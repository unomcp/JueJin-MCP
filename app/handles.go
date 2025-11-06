package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

func (app *App) healthHandle(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func (app *App) mcpHandle() fiber.Handler {
	mcpHandler := goMcp.NewStreamableHTTPHandler(func(r *http.Request) *goMcp.Server {
		return app.MCP
	}, &goMcp.StreamableHTTPOptions{
		JSONResponse: true,
	})
	return adaptor.HTTPHandler(mcpHandler)
}

func (app *App) loginStatusHandle(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
