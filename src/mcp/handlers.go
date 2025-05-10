package mcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"kikokai/src/model"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
)

// Event types for SSE (copied from server.go to maintain consistency)
type CubeEvent struct {
	Type      string                `json:"type"`
	Axis      string                `json:"axis,omitempty"`      // x, y, z
	Layer     int                   `json:"layer,omitempty"`     // 1 ou -1
	Direction int                   `json:"direction,omitempty"` // 1 pour sens horaire, -1 pour sens anti-horaire
	State     [3][3][3]*model.Cubie `json:"state,omitempty"`
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

// Adapted to the new cube structure
func stateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("Received MCP request: %s", CommandState)

	// Return the cube state using the updated structure
	return mcp.NewToolResultText(fmt.Sprintf("Cube state: %v", model.SharedCube.Cubies)), nil
}

func resetHandler(tx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("Received MCP request: %s", CommandReset)

	// Broadcast the reset event
	if Broadcaster != nil {
		Broadcaster.BroadcastEvent(CubeEvent{
			Type: "reset",
		})
	}

	// Reset the cube using the new structure
	model.ResetCube()

	// Send the response
	return mcp.NewToolResultText(fmt.Sprintf("Cube state: %v", model.SharedCube.Cubies)), nil
}

func scrambleHandler(tx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("Received MCP request: %s", CommandScramble)

	// Broadcast the scramble event
	if Broadcaster != nil {
		Broadcaster.BroadcastEvent(CubeEvent{
			Type: "scramble",
		})
	}

	// Scramble the cube using the new structure
	model.SharedCube.Scramble(20) // Scramble with 20 random moves

	// Send the response
	return mcp.NewToolResultText(fmt.Sprintf("Cube state: %v", model.SharedCube.Cubies)), nil
}

func rotateAxisHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Printf("Received MCP request: rotate_axis")

	// Extract parameters from the request
	axis, ok := request.Params.Arguments["axis"].(string)
	if !ok {
		return nil, errors.New("axis must be a string ('x', 'y', or 'z')")
	}

	layer, err := getFloatParam(request.Params.Arguments, "layer")
	if err != nil {
		return nil, err
	}

	direction, err := getFloatParam(request.Params.Arguments, "direction")
	if err != nil {
		return nil, err
	}

	// Validate inputs
	if axis != "x" && axis != "y" && axis != "z" {
		return nil, errors.New("axis must be 'x', 'y', or 'z'")
	}

	if layer != 1 && layer != -1 {
		return nil, errors.New("layer must be 1 or -1")
	}

	if direction != 1 && direction != -1 {
		return nil, errors.New("direction must be 1 (clockwise) or -1 (counter-clockwise)")
	}

	// Map axis, layer, and direction to face and rotation direction
	face := model.GetCoordFromAxis(axis, int(layer))
	clockwise := model.TurningDirection(direction == 1)

	// Broadcast the rotation event
	if Broadcaster != nil {
		log.Printf("Broadcasting MCP axis rotation: axis=%s, layer=%d, direction=%d", axis, int(layer), int(direction))
		Broadcaster.BroadcastEvent(CubeEvent{
			Type:      "rotate",
			Axis:      axis,
			Layer:     int(layer),
			Direction: int(direction),
		})
	}

	// Apply the rotation to the cube
	model.SharedCube.RotateAxis(face, clockwise)

	// Send the response
	return mcp.NewToolResultText(fmt.Sprintf("Rotated cube: axis=%s, layer=%d, direction=%d", axis, int(layer), int(direction))), nil
}

// Fonction utilitaire pour extraire un paramètre numérique
func getFloatParam(args map[string]interface{}, name string) (float64, error) {
	val, ok := args[name]
	if !ok {
		return 0, fmt.Errorf("%s parameter is required", name)
	}

	// MCP peut renvoyer un nombre sous forme de float64 ou json.Number
	switch v := val.(type) {
	case float64:
		return v, nil
	case int:
		return float64(v), nil
	case json.Number:
		f, err := v.Float64()
		if err != nil {
			return 0, fmt.Errorf("invalid %s value: %v", name, err)
		}
		return f, nil
	default:
		return 0, fmt.Errorf("%s must be a number", name)
	}
}
