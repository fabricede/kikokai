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
	State [6][3][3]string `json:"state"`
}

func main() {
	// Setup routes
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/api/state", getStateHandler)
	http.HandleFunc("/api/rotate", rotateHandler)
	http.HandleFunc("/api/reset", resetHandler)

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
	clockwise := model.Direction(req.Clockwise)

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
