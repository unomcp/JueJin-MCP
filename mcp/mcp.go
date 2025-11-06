package mcp

import (
	"context"

	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct {
	Name string `json:"name" jsonschema:"the name of the person to greet"`
}

type Output struct {
	Greeting string `json:"greeting" jsonschema:"the greeting to tell to the user"`
}

func SayHi(ctx context.Context, req *goMcp.CallToolRequest, input Input) (
	*goMcp.CallToolResult,
	Output,
	error,
) {
	return nil, Output{Greeting: "Hi " + input.Name}, nil
}

func InitMCP() *goMcp.Server {
	server := goMcp.NewServer(&goMcp.Implementation{
		Name:    "JueJin-MCP",
		Version: "0.0.1",
	}, nil)

	// 添加工具
	goMcp.AddTool(server, &goMcp.Tool{
		Name:        "greet",
		Description: "say hi",
	}, SayHi)

	// 添加登录状态工具
	goMcp.AddTool(server, &goMcp.Tool{
		Name:        "loginStatus",
		Description: "get login status",
	}, LoginStatus)
	return server
}
