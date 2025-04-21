package main

import (
	"encoding/json"
	"fmt"
	"kikokai/src/mcp"
	"kikokai/src/model"
	"kikokai/src/shared"
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

type CubeStateResponse struct {
	State [6]model.Face `json:"state"`
}

// Event types for SSE
type CubeEvent struct {
	Type      string        `json:"type"`
	Face      int           `json:"face,omitempty"`
	Clockwise bool          `json:"clockwise,omitempty"`
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
		State: shared.Cube.State,
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

func main() {
	// Set the correct MIME type for WebAssembly files
	mime.AddExtensionType(".wasm", "application/wasm")

	// Start the event broker
	broker.Start()

	// Setup routes
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/api/state", getStateHandler)
	http.HandleFunc("/api/rotate", rotateHandler)
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
	json.NewEncoder(w).Encode(CubeStateResponse{State: shared.Cube.State})
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
	shared.Cube.RotateFace(face, clockwise)

	// Send the response with updated state
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CubeStateResponse{State: shared.Cube.State})
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to reset cube")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	shared.ResetCube()

	// Broadcast the reset event
	broker.BroadcastEvent(CubeEvent{
		Type: "reset",
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CubeStateResponse{State: shared.Cube.State})
}

// scrambleHandler randomly scrambles the cube
func scrambleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to scramble cube")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Use the new Scramble method from the refactored cube model
	shared.Cube.Scramble(20) // Scramble with 20 random moves

	// Broadcast the scramble event
	broker.BroadcastEvent(CubeEvent{
		Type: "scramble",
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CubeStateResponse{State: shared.Cube.State})
}
