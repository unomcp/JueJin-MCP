package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

func SetupRoutes(app *App) {
	mcpHandler := goMcp.NewStreamableHTTPHandler(func(r *http.Request) *goMcp.Server {
		return app.MCP
	}, &goMcp.StreamableHTTPOptions{
		JSONResponse: true,
	})

	// 探活接口
	app.FiberApp.All("/health", adaptor.HTTPHandler(mcpHandler))
	// MCP 接口
	app.FiberApp.All("/mcp", adaptor.HTTPHandler(mcpHandler))
}
