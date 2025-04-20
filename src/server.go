package main

import (
	"encoding/json"
	"kikokai/src/mcp"
	"kikokai/src/model"
	"kikokai/src/shared"
	"log"
	"net/http"
)

// API Request and Response types
type RotateRequest struct {
	Face      int  `json:"face"`
	Clockwise bool `json:"clockwise"`
}

type CubeStateResponse struct {
	State [6]model.Face `json:"state"`
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

	// Use the new Scramble method from the refactored cube model
	shared.Cube.Scramble(20) // Scramble with 20 random moves

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CubeStateResponse{State: shared.Cube.State})
}
