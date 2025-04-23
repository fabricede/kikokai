package main

import (
	"encoding/json"
	"fmt"
	"kikokai/src/mcp"
	"kikokai/src/model"
	"log"
	"mime"
	"net/http"
	"sync"
)

// API Request and Response types
type RotateRequest struct {
	Face      int  `json:"face"`
	Clockwise bool `json:"clockwise"`
}

// New request structure for axis-based rotations
type RotateAxisRequest struct {
	Axis      string `json:"axis"`      // "x", "y", or "z"
	Layer     int    `json:"layer"`     // 1 or -1
	Direction int    `json:"direction"` // 1 for clockwise, -1 for counter-clockwise
}

type CubeStateResponse struct {
	State [6]model.Face `json:"state"`
}

// Event types for SSE
type CubeEvent struct {
	Type      string        `json:"type"`
	Face      int           `json:"face,omitempty"`
	Clockwise bool          `json:"clockwise,omitempty"`
	Axis      string        `json:"axis,omitempty"`      // x, y, z
	Layer     int           `json:"layer,omitempty"`     // 1 ou -1
	Direction int           `json:"direction,omitempty"` // 1 pour sens horaire, -1 pour sens anti-horaire
	State     [6]model.Face `json:"state,omitempty"`
}

// EventBroker manages SSE connections
type EventBroker struct {
	// Registered clients
	clients map[chan []byte]bool

	// Register requests
	register chan chan []byte

	// Unregister requests
	unregister chan chan []byte

	// Events to broadcast to clients
	broadcast chan []byte

	// Mutex for thread safety
	mutex sync.Mutex
}

// Create a new event broker
func NewEventBroker() *EventBroker {
	return &EventBroker{
		clients:    make(map[chan []byte]bool),
		register:   make(chan chan []byte),
		unregister: make(chan chan []byte),
		broadcast:  make(chan []byte, 10),
	}
}

// Start the event broker
func (eb *EventBroker) Start() {
	go func() {
		for {
			select {
			case client := <-eb.register:
				eb.mutex.Lock()
				eb.clients[client] = true
				eb.mutex.Unlock()
				log.Println("Client registered for event updates")

			case client := <-eb.unregister:
				eb.mutex.Lock()
				if _, ok := eb.clients[client]; ok {
					delete(eb.clients, client)
					close(client)
				}
				eb.mutex.Unlock()
				log.Println("Client unregistered from event updates")

			case message := <-eb.broadcast:
				eb.mutex.Lock()
				for client := range eb.clients {
					select {
					case client <- message:
					default:
						close(client)
						delete(eb.clients, client)
					}
				}
				eb.mutex.Unlock()
			}
		}
	}()
}

// ServeHTTP implements the http.Handler interface
func (eb *EventBroker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create a channel for this client
	messageChan := make(chan []byte)

	// Register this client
	eb.register <- messageChan

	// Remove client when connection closes
	go func() {
		// Use the request's context for cancellation
		<-r.Context().Done()
		eb.unregister <- messageChan
	}()

	// Send initial state event
	initialState, _ := json.Marshal(CubeEvent{
		Type:  "state",
		State: model.SharedCube.State,
	})
	fmt.Fprintf(w, "data: %s\n\n", initialState)
	w.(http.Flusher).Flush()

	// Stream events to client
	for {
		msg, open := <-messageChan
		if !open {
			break
		}

		fmt.Fprintf(w, "data: %s\n\n", msg)
		w.(http.Flusher).Flush()
	}
}

// Broadcast an event to all connected clients
func (eb *EventBroker) BroadcastEvent(event interface{}) {
	data, err := json.Marshal(event)
	if err != nil {
		log.Println("Error marshalling event:", err)
		return
	}
	eb.broadcast <- data
}

// Create a global event broker
var broker = NewEventBroker()

// Connect the MCP broadcaster to our broker
func init() {
	// Set the MCP broadcaster to use our event broker
	mcp.Broadcaster = broker
}

func main() {
	// Set the correct MIME type for WebAssembly files
	mime.AddExtensionType(".wasm", "application/wasm")

	// Start the event broker
	broker.Start()

	// Setup routes
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/api/state", getStateHandler)
	http.HandleFunc("/api/rotate", rotateHandler)
	http.HandleFunc("/api/rotate-axis", rotateAxisHandler) // Nouvelle route pour la rotation par axe
	http.HandleFunc("/api/reset", resetHandler)
	http.HandleFunc("/api/scramble", scrambleHandler)
	http.Handle("/api/events", broker)

	// Start MCP server in a goroutine
	go mcp.StartMCPServer()

	// Start the HTTP server
	log.Println("HTTP server started at http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func getStateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request for cube state")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CubeStateResponse{State: model.SharedCube.State})
}

func rotateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to rotate cube face")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RotateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	face := model.FaceIndex(req.Face)
	// Convert boolean to Direction type
	var clockwise model.TurningDirection
	if req.Clockwise {
		clockwise = model.Clockwise
	} else {
		clockwise = model.CounterClockwise
	}

	if face < 0 || face > 5 {
		http.Error(w, "Invalid face index", http.StatusBadRequest)
		return
	}

	// Log the exact values being sent to ensure they're correct
	log.Printf("Broadcasting rotation event: face=%d, clockwise=%t", req.Face, req.Clockwise)

	// Broadcast the rotation event to connected clients
	broker.BroadcastEvent(CubeEvent{
		Type:      "rotate",
		Face:      req.Face,
		Clockwise: req.Clockwise,
	})

	// Apply rotation to the server-side cube after broadcasting event
	// This ensures the animation happens first, then state is updated
	model.SharedCube.RotateFace(face, clockwise)

	// Send the response with updated state
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CubeStateResponse{State: model.SharedCube.State})
}

// rotateAxisHandler traite les rotations basées sur l'axe, la couche et la direction
func rotateAxisHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to rotate cube by axis")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RotateAxisRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validation des entrées
	if req.Axis != "x" && req.Axis != "y" && req.Axis != "z" {
		http.Error(w, "Invalid axis, must be 'x', 'y', or 'z'", http.StatusBadRequest)
		return
	}

	if req.Layer != 1 && req.Layer != -1 {
		http.Error(w, "Invalid layer, must be 1 or -1", http.StatusBadRequest)
		return
	}

	if req.Direction != 1 && req.Direction != -1 {
		http.Error(w, "Invalid direction, must be 1 (clockwise) or -1 (counter-clockwise)", http.StatusBadRequest)
		return
	}

	// Table de conversion de l'axe, couche et direction vers face et sens de rotation
	// Le sens de rotation est inversé pour certaines faces pour maintenir une cohérence dans le modèle 3D
	var face model.FaceIndex
	var clockwise model.TurningDirection

	// Définir la face et le sens de rotation en fonction de l'axe, de la couche et de la direction
	switch req.Axis {
	case "x":
		if req.Layer == 1 { // Front face (x=1)
			face = model.Front
			clockwise = model.TurningDirection(req.Direction == 1) // 1 = Clockwise, -1 = CounterClockwise
		} else { // Back face (x=-1)
			face = model.Back
			clockwise = model.TurningDirection(req.Direction == -1) // Sens inversé pour la face arrière
		}
	case "y":
		if req.Layer == 1 { // Up face (y=1)
			face = model.Up
			clockwise = model.TurningDirection(req.Direction == 1)
		} else { // Down face (y=-1)
			face = model.Down
			clockwise = model.TurningDirection(req.Direction == -1) // Sens inversé pour la face inférieure
		}
	case "z":
		if req.Layer == 1 { // Right face (z=1)
			face = model.Right
			clockwise = model.TurningDirection(req.Direction == 1)
		} else { // Left face (z=-1)
			face = model.Left
			clockwise = model.TurningDirection(req.Direction == -1) // Sens inversé pour la face gauche
		}
	}

	// Journalisation de la conversion
	log.Printf("Axis rotation: axis=%s, layer=%d, direction=%d mapped to face=%d, clockwise=%t",
		req.Axis, req.Layer, req.Direction, face, clockwise)

	// Diffuser l'événement de rotation aux clients connectés avec les paramètres d'axe originaux
	broker.BroadcastEvent(CubeEvent{
		Type:      "rotate",
		Axis:      req.Axis,
		Layer:     req.Layer,
		Direction: req.Direction,
	})

	// Appliquer la rotation au cube côté serveur
	model.SharedCube.RotateFace(face, clockwise)

	// Envoyer la réponse avec l'état mis à jour
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CubeStateResponse{State: model.SharedCube.State})
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to reset cube")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	model.ResetCube()

	// Broadcast the reset event
	broker.BroadcastEvent(CubeEvent{
		Type: "reset",
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CubeStateResponse{State: model.SharedCube.State})
}

// scrambleHandler randomly scrambles the cube
func scrambleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to scramble cube")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Use the new Scramble method from the refactored cube model
	model.SharedCube.Scramble(20) // Scramble with 20 random moves

	// Broadcast the scramble event
	broker.BroadcastEvent(CubeEvent{
		Type: "scramble",
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CubeStateResponse{State: model.SharedCube.State})
}
