package mcp

import (
	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/unomcp/JueJin-MCP/juejin"
)

type MCP struct {
	JueJin *juejin.JueJin
}

func NewMCP(jj *juejin.JueJin) *MCP {
	return &MCP{
		JueJin: jj,
	}
}

func InitMCP(jj *juejin.JueJin) *goMcp.Server {
	mcpInstance := NewMCP(jj)

	server := goMcp.NewServer(&goMcp.Implementation{
		Name:    "JueJin-MCP",
		Version: "0.0.1",
	}, nil)

	// 添加登录状态工具
	goMcp.AddTool(server, &goMcp.Tool{
		Name:        "login status",
		Description: "获取登录状态",
	}, mcpInstance.LoginStatus)

	// 添加发布工具
	goMcp.AddTool(server, &goMcp.Tool{
		Name:        "publish article",
		Description: "发布文章",
	}, mcpInstance.PublishArticle)

	return server
}
