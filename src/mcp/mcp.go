package mcp

import (
	"encoding/json"
	"kikokai/src/model"
	"kikokai/src/shared"
	"log"
	"net"
)

// MCP Protocol structures
type MCPRequest struct {
	Command string `json:"command"`
	Params  struct {
		Face      int  `json:"face"`
		Clockwise bool `json:"clockwise"`
	} `json:"params"`
}

type MCPResponse struct {
	State [6][3][3]string `json:"state"`
	Error string          `json:"error,omitempty"`
}

// MCP Command constants
const (
	CommandRotate = "rotate"
	CommandReset  = "reset"
	CommandState  = "state"
)

func handleMCPConnection(conn net.Conn) {
	defer conn.Close()

	// Create a decoder for the connection
	decoder := json.NewDecoder(conn)

	// Read the request
	var req MCPRequest
	if err := decoder.Decode(&req); err != nil {
		log.Println("Failed to decode request:", err)
		sendErrorResponse(conn, "Invalid request format")
		return
	}

	log.Printf("Received MCP request: %s", req.Command)

	// Process the request based on the command
	var resp MCPResponse
	resp.State = shared.Cube.State

	switch req.Command {
	case CommandRotate:
		face := model.Face(req.Params.Face)
		direction := model.Direction(req.Params.Clockwise)

		if face < 0 || face > 5 {
			resp.Error = "Invalid face index"
		} else {
			shared.Cube.RotateFace(face, direction)
			resp.State = shared.Cube.State
		}

	case CommandReset:
		shared.ResetCube()
		resp.State = shared.Cube.State

	case CommandState:
		// Just return the current state

	default:
		resp.Error = "Unknown command"
	}

	// Send the response
	if err := json.NewEncoder(conn).Encode(resp); err != nil {
		log.Println("Failed to encode response:", err)
	}
}

func sendErrorResponse(conn net.Conn, errMsg string) {
	resp := MCPResponse{
		Error: errMsg,
	}
	json.NewEncoder(conn).Encode(resp)
}

// StartMCPServer starts a TCP server for MCP communication
func StartMCPServer() {
	ln, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal("Failed to start MCP server:", err)
	}
	defer ln.Close()
	log.Println("MCP server started on :9001")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}
		go handleMCPConnection(conn)
	}
}
