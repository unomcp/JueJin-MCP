package mcp

import goMcp "github.com/modelcontextprotocol/go-sdk/mcp"

func InitMCP() {
	goMcp.NewServer(&goMcp.Implementation{
		Name:    "JueJin-MCP",
		Title:   "JueJin-MCP",
		Version: "0.0.1",
	}, nil)
}
