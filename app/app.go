package app

import (
	"github.com/gofiber/fiber/v2"
	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/unomcp/JueJin-MCP/juejin"
	"github.com/unomcp/JueJin-MCP/mcp"
)

type App struct {
	JueJin   *juejin.JueJin
	MCP      *goMcp.Server
	FiberApp *fiber.App
}

func NewApp(app *fiber.App) *App {
	return &App{
		JueJin:   juejin.NewJueJin(),
		MCP:      mcp.InitMCP(),
		FiberApp: app,
	}
}
