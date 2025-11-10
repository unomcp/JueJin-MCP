package mcp

import (
	"context"

	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

func (m *MCP) LoginStatus(ctx context.Context, req *goMcp.CallToolRequest, _ any) (
	*goMcp.CallToolResult,
	any,
	error,
) {
	// 现在可以通过 m.JueJin 调用 juejin 中的方法
	// 例如: m.JueJin.CheckLoginStatus()
	m.JueJin.Login()

	return nil, nil, nil
}

func (m *MCP) PublishArticle(ctx context.Context, req *goMcp.CallToolRequest, _ any) (
	*goMcp.CallToolResult,
	any,
	error,
) {
	m.JueJin.PublishArticle()

	return nil, nil, nil
}
