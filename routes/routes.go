package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/unomcp/JueJin-MCP/mcp"
)

func SetupRoutes(app *fiber.App) {
	mcpHandler := goMcp.NewStreamableHTTPHandler(func(r *http.Request) *goMcp.Server {
		return mcp.InitMCP()
	}, &goMcp.StreamableHTTPOptions{
		JSONResponse: true,
	})
	app.Get("/mcp", adaptor.HTTPHandler(mcpHandler))
}
