package main

import (
	"encoding/json"
	"kikokai/src/mcp"
	"kikokai/src/model"
	"kikokai/src/shared"
	"log"
	"math/rand"
	"net/http"
)

// API Request and Response types
type RotateRequest struct {
	Face      int  `json:"face"`
	Clockwise bool `json:"clockwise"`
}

type CubeStateResponse struct {
	State [6][3][3]model.Color `json:"state"`
}

func main() {
	// Setup routes
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/api/state", getStateHandler)
	http.HandleFunc("/api/rotate", rotateHandler)
	http.HandleFunc("/api/reset", resetHandler)
	http.HandleFunc("/api/scramble", scrambleHandler)

	// Start MCP server in a goroutine
	go mcp.StartMCPServer()

	// Start the HTTP server
	log.Println("HTTP server started at http://localhost:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func getStateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CubeStateResponse{State: shared.Cube.State})
}

func rotateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RotateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	face := model.Face(req.Face)
	clockwise := model.TurningDirection(req.Clockwise)

	if face < 0 || face > 5 {
		http.Error(w, "Invalid face index", http.StatusBadRequest)
		return
	}

	shared.Cube.RotateFace(face, clockwise)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CubeStateResponse{State: shared.Cube.State})
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	shared.ResetCube()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CubeStateResponse{State: shared.Cube.State})
}

// scrambleHandler randomly scrambles the cube
func scrambleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Scramble the cube with 20-25 random moves
	numMoves := 20 + rand.Intn(6) // Random number between 20 and 25
	log.Printf("Scrambling cube with %d random moves", numMoves)

	for i := 0; i < numMoves; i++ {
		face := model.Face(rand.Intn(6))                       // Random face (0-5)
		clockwise := model.TurningDirection(rand.Intn(2) == 1) // Random direction
		shared.Cube.RotateFace(face, clockwise)
		log.Printf("Scramble move %d: Face %d, Clockwise: %v", i+1, face, clockwise)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CubeStateResponse{State: shared.Cube.State})
}
