package mcp

import (
	"context"
	"log"

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

	log.Println("MCP Server initialized with greet tool")

	// HTTP 模式下不需要调用 server.Run()
	// server.Run() 只用于 stdio 模式
	return server
}
