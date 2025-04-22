package mcp

import (
	"context"
	"errors"
	"fmt"
	"kikokai/src/model"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
)

// Event types for SSE (copied from server.go to maintain consistency)
type CubeEvent struct {
	Type      string        `json:"type"`
	Face      int           `json:"face,omitempty"`
	Clockwise bool          `json:"clockwise,omitempty"`
	State     [6]model.Face `json:"state,omitempty"`
}

// Interface for broadcasting events
type EventBroadcaster interface {
	BroadcastEvent(event any)
}

// Global broadcaster that will be set by the main package
var Broadcaster EventBroadcaster

// MCP Command constants
const (
	CommandRotate   = "rotate"
	CommandReset    = "reset"
	CommandState    = "state"
	CommandScramble = "scramble"
)

func stateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("Received MCP request: %s", CommandState)

	return mcp.NewToolResultText(fmt.Sprintf("Cube state: %v", model.SharedCube.State)), nil
}

func rotateHandler(tx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("Received MCP request: %s", CommandRotate)
	face, ok := request.Params.Arguments["face"].(int)
	if !ok {
		return nil, errors.New("face must be an int")
	}
	clockwise, ok := request.Params.Arguments["clockwise"].(bool)
	if !ok {
		return nil, errors.New("clockwise must be a bool")
	}

	if face < 0 || face > 5 {
		return nil, errors.New("Invalid face index")
	} else {
		// Broadcast the rotation event to browser clients if broadcaster is set
		if Broadcaster != nil {
			log.Printf("Broadcasting MCP rotation event: face=%d, clockwise=%t", face, clockwise)
			Broadcaster.BroadcastEvent(CubeEvent{
				Type:      "rotate",
				Face:      face,
				Clockwise: clockwise,
			})
		}

		model.SharedCube.RotateFace(model.FaceIndex(face), model.TurningDirection(clockwise))
	}

	// Send the response
	return mcp.NewToolResultText(fmt.Sprintf("Cube state: %v", model.SharedCube.State)), nil
}

func resetHandler(tx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("Received MCP request: %s", CommandReset)

	// Broadcast the reset event
	if Broadcaster != nil {
		Broadcaster.BroadcastEvent(CubeEvent{
			Type: "reset",
		})
	}

	model.ResetCube()
	// Send the response
	return mcp.NewToolResultText(fmt.Sprintf("Cube state: %v", model.SharedCube.State)), nil
}

func scrambleHandler(tx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("Received MCP request: %s", CommandScramble)

	// Broadcast the scramble event
	if Broadcaster != nil {
		Broadcaster.BroadcastEvent(CubeEvent{
			Type: "scramble",
		})
	}

	// Add scramble support to the MCP server
	model.SharedCube.Scramble(20) // Scramble with 20 random moves
	// Send the response
	return mcp.NewToolResultText(fmt.Sprintf("Cube state: %v", model.SharedCube.State)), nil
}
