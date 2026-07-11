package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"worldcup-injective/internal/matches"
	"worldcup-injective/internal/predictions"
)

// StartMCPServer starts an MCP server over stdio so AI agents / MCP clients
// can query World Cup matches and predictions. This is one of the required
// Injective technologies for the hackathon (MCP Server + Agent Skills).
func StartMCPServer() {
	s := server.NewMCPServer(
		"worldcup-injective",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	// Tool: list all World Cup matches.
	listTool := mcp.NewTool("list_worldcup_matches",
		mcp.WithDescription("List all available World Cup matches with teams, odds and status."),
	)
	s.AddTool(listTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		data, _ := json.MarshalIndent(matches.All(), "", "  ")
		return mcp.NewToolResultText(string(data)), nil
	})

	// Tool: predict a specific match.
	predictTool := mcp.NewTool("predict_match",
		mcp.WithDescription("Predict the outcome of a World Cup match by its ID."),
		mcp.WithString("match_id",
			mcp.Required(),
			mcp.Description("The match ID, e.g. 'final-arg-fra'."),
		),
	)
	s.AddTool(predictTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		id, err := req.RequireString("match_id")
		if err != nil {
			return mcp.NewToolResultError("match_id is required"), nil
		}
		m, ok := matches.ByID(id)
		if !ok {
			return mcp.NewToolResultError(fmt.Sprintf("match %q not found", id)), nil
		}
		pred := predictions.Premium(m)
		data, _ := json.MarshalIndent(pred, "", "  ")
		return mcp.NewToolResultText(string(data)), nil
	})

	log.Println("MCP server started (stdio). Tools: list_worldcup_matches, predict_match")
	if err := server.ServeStdio(s); err != nil {
		log.Printf("MCP server stopped: %v", err)
	}
}
