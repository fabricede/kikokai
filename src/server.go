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

// Request structure for axis-based rotations
type RotateAxisRequest struct {
	Axis      string `json:"axis"`      // "x", "y", or "z"
	Layer     int    `json:"layer"`     // 1 or -1
	Direction int    `json:"direction"` // 1 for clockwise, -1 for counter-clockwise
}

type CubeStateResponse struct {
	State [3][3][3]*model.Cubie `json:"state"`
}

// Event types for SSE
type CubeEvent struct {
	Type      string                `json:"type"`
	Axis      string                `json:"axis,omitempty"`      // x, y, z
	Layer     int                   `json:"layer,omitempty"`     // 1 ou -1
	Direction int                   `json:"direction,omitempty"` // 1 pour sens horaire, -1 pour sens anti-horaire
	State     [3][3][3]*model.Cubie `json:"state,omitempty"`
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
		State: model.SharedCube.Cubies,
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
	http.HandleFunc("/api/state", handleState)
	http.HandleFunc("/api/rotate-axis", handleRotate) // Nouvelle route pour la rotation par axe
	http.HandleFunc("/api/reset", handleReset)
	http.HandleFunc("/api/scramble", handleScramble)
	http.Handle("/api/events", broker)

	// Start MCP server in a goroutine
	go mcp.StartMCPServer()

	// Start the HTTP server
	log.Println("HTTP server started at http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}

// Adapted to the new cube structure
func handleState(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling state request")

	// Return the cube state using the updated structure
	response := struct {
		State [3][3][3]*model.Cubie `json:"state"`
	}{
		State: model.SharedCube.Cubies,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding state response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling reset request")

	// Reset the cube using the new structure
	model.ResetCube()

	// Broadcast the reset event
	broker.BroadcastEvent(CubeEvent{
		Type: "reset",
	})

	// Return the updated state
	handleState(w, r)
}

func handleScramble(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling scramble request")

	// Scramble the cube using the new structure
	model.SharedCube.Scramble(20) // Scramble with 20 random moves

	// Broadcast the scramble event
	broker.BroadcastEvent(CubeEvent{
		Type: "scramble",
	})

	// Return the updated state
	handleState(w, r)
}

func handleRotate(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling rotate request")

	// Parse the request body
	var req struct {
		Axis      string `json:"axis"`
		Layer     int    `json:"layer"`
		Direction int    `json:"direction"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding rotate request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate inputs
	if req.Axis != "x" && req.Axis != "y" && req.Axis != "z" {
		http.Error(w, "Invalid axis; must be 'x', 'y', or 'z'", http.StatusBadRequest)
		return
	}
	if req.Layer != 1 && req.Layer != -1 {
		http.Error(w, "Invalid layer; must be 1 or -1", http.StatusBadRequest)
		return
	}
	if req.Direction != 1 && req.Direction != -1 {
		http.Error(w, "Invalid direction; must be 1 or -1", http.StatusBadRequest)
		return
	}

	// Map axis, layer, and direction to face and rotation direction
	var face model.FaceIndex
	var clockwise model.TurningDirection

	switch req.Axis {
	case "x":
		if req.Layer == 1 {
			face = model.Front
			clockwise = model.TurningDirection(req.Direction == 1)
		} else {
			face = model.Back
			clockwise = model.TurningDirection(req.Direction == -1)
		}
	case "y":
		if req.Layer == 1 {
			face = model.Up
			clockwise = model.TurningDirection(req.Direction == 1)
		} else {
			face = model.Down
			clockwise = model.TurningDirection(req.Direction == -1)
		}
	case "z":
		if req.Layer == 1 {
			face = model.Right
			clockwise = model.TurningDirection(req.Direction == 1)
		} else {
			face = model.Left
			clockwise = model.TurningDirection(req.Direction == -1)
		}
	}

	// Apply the rotation to the cube
	model.SharedCube.RotateFace(face, clockwise)

	// Broadcast the rotation event
	broker.BroadcastEvent(CubeEvent{
		Type:      "rotate",
		Axis:      req.Axis,
		Layer:     req.Layer,
		Direction: req.Direction,
	})

	// Return the updated state
	handleState(w, r)
}
