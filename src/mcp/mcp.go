package mcp

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const describes = `Send http API request to interact with a rubick's cube, 
possible action are 
 - 'state' to retreive the current state of the cube, 
 - 'reset' to return to initial value, 
 - 'scramble' to scramble randomly
 - 'rotate' to rotate a cube layer based on axis, layer and direction this action requires a body to indicate the axis (x, y, z), layer (1 or -1) and direction (1 for clockwise, -1 for counter-clockwise)"),
`

func StartMCPServer() {
	// Create MCP server
	mcpServer := server.NewMCPServer(
		"Demo Rubic's cube 🚀",
		"1.0.0",
	)

	httpTool := mcp.NewTool("http_request",
		mcp.WithDescription("Send http API request to interact with a rubick's cube"),
		mcp.WithString("url",
			mcp.Required(),
			mcp.Description("URL of the cube API "),
			mcp.Pattern("^http:localhost:8090//api/.*"),
		),
		mcp.WithString("method",
			mcp.Required(),
			mcp.Description("HTTP method (POST)"),
		),
		mcp.WithString("body",
			mcp.Description("Request body"),
		),
	)

	// Add tool
	mcpServer.AddTool(httpTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Extract parameters
		url := request.Params.Arguments["url"].(string)
		method := request.Params.Arguments["method"].(string)
		body := ""
		if b, ok := request.Params.Arguments["body"].(string); ok {
			body = b
		}
		log.Printf("Received MCP request: %s", url)
		log.Printf("Method: %s", method)
		log.Printf("Body: %s", body)
		// Call the HTTP API
		// Create and send request
		var req *http.Request
		var err error
		if body != "" {
			req, err = http.NewRequest(method, url, strings.NewReader(body))
		} else {
			req, err = http.NewRequest(method, url, nil)
		}
		if err != nil {
			return mcp.NewToolResultErrorFromErr("unable to create request", err), nil
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("unable to execute request", err), nil
		}
		defer resp.Body.Close()

		// Return response
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("unable to read request response", err), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Status: %d\nBody: %s", resp.StatusCode, string(respBody))), nil
	})

	// Add state tool
	getState := mcp.NewTool("state",
		mcp.WithDescription("get the state of the cube"),
	)
	// Add tool handler
	mcpServer.AddTool(getState, stateHandler)

	// Add rotate-axis tool
	rotateAxis := mcp.NewTool("rotate_axis",
		mcp.WithDescription("rotate a cube layer based on axis, layer and direction"),
		mcp.WithString("axis",
			mcp.Required(),
			mcp.Description("Rotation axis (x, y, z)"),
		),
		mcp.WithNumber("layer",
			mcp.Required(),
			mcp.Description("Layer to rotate (1 or -1)"),
		),
		mcp.WithNumber("direction",
			mcp.Required(),
			mcp.Description("Rotation direction (1 for clockwise, -1 for counter-clockwise)"),
		),
	)
	// Add rotate-axis tool handler
	mcpServer.AddTool(rotateAxis, rotateAxisHandler)

	// Add reset tool
	reset := mcp.NewTool("reset",
		mcp.WithDescription("reset the cube"),
	)
	// Add reset tool handler
	mcpServer.AddTool(reset, resetHandler)

	// Add scramble tool
	scramble := mcp.NewTool("scramble",
		mcp.WithDescription("scramble the cube"),
	)
	// Add scramble tool handler
	mcpServer.AddTool(scramble, scrambleHandler)

	// Starts the sse server
	sseServer := server.NewSSEServer(mcpServer, server.WithBaseURL("http://localhost:9001"))
	log.Printf("SSE server listening on :9001")
	if err := sseServer.Start(":9001"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
