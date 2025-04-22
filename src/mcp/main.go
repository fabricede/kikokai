package mcp

import (
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func StartMCPServer() {
	log.Println("Starting MCP server...")
	// Create MCP server
	s := server.NewMCPServer(
		"Demo Rubic's cube 🚀",
		"1.0.0",
	)

	// Add tool
	getState := mcp.NewTool("state",
		mcp.WithDescription("get the state of the cube"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the cube"),
		),
	)
	// Add tool handler
	s.AddTool(getState, stateHandler)

	// Add rotate tool
	rotate := mcp.NewTool("rotate",
		mcp.WithDescription("rotate a face of the cube"),
		mcp.WithNumber("face",
			mcp.Required(),
			mcp.Description("Face index (0-5)"),
		),
		mcp.WithBoolean("clockwise",
			mcp.Required(),
			mcp.Description("Rotate clockwise"),
		),
	)
	// Add rotate tool handler
	s.AddTool(rotate, rotateHandler)

	// Add reset tool
	reset := mcp.NewTool("reset",
		mcp.WithDescription("reset the cube"),
	)
	// Add reset tool handler
	s.AddTool(reset, resetHandler)

	// Add scramble tool
	scramble := mcp.NewTool("scramble",
		mcp.WithDescription("scramble the cube"),
	)
	// Add scramble tool handler
	s.AddTool(scramble, scrambleHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
