package mcp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/unomcp/JueJin-MCP/configs"
)

func initMCP() *goMcp.Server {
	server := goMcp.NewServer(&goMcp.Implementation{
		Name:    configs.MCPName,
		Version: configs.MCPVersion,
	}, nil)

	// 添加登录工具
	goMcp.AddTool(server, &goMcp.Tool{
		Name:        "login",
		Description: "登录掘金",
	}, loginTool)

	// 添加草稿工具
	goMcp.AddTool(server, &goMcp.Tool{
		Name:        "draft",
		Description: "写掘金文章保存到草稿箱（不发布）",
	}, draftTool)

	// 添加发布工具
	goMcp.AddTool(server, &goMcp.Tool{
		Name:        "publish",
		Description: "发布掘金文章",
	}, publishTool)

	return server
}

func MCP() fiber.Handler {
	mcp := initMCP()
	mcpHandler := goMcp.NewStreamableHTTPHandler(func(r *http.Request) *goMcp.Server {
		return mcp
	}, configs.MCPStreamableHTTPOptions)
	return adaptor.HTTPHandler(mcpHandler)
}
