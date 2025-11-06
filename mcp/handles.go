package mcp

import (
	"context"

	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

func LoginStatus(ctx context.Context, req *goMcp.CallToolRequest, input Input) (
	*goMcp.CallToolResult,
	Output,
	error,
) {
	return nil, Output{Greeting: "Login Status"}, nil
}
