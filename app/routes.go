package app

import (
	"github.com/unomcp/JueJin-MCP/middleware"
)

func (app *App) SetupRoutes() {
	app.FiberApp.Use(middleware.Cros())

	// 探活接口
	app.FiberApp.All("/health", app.healthHandle)
	// MCP 接口
	app.FiberApp.All("/mcp", app.mcpHandle())

	app.FiberApp.Get("/login/status", app.loginStatusHandle)
}
