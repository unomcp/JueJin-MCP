package juejin

import (
	"log"

	"github.com/gofiber/fiber/v2"
	goMcp "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/unomcp/JueJin-MCP/mcp"
	"github.com/unomcp/JueJin-MCP/middleware"
	"github.com/unomcp/JueJin-MCP/routes"
)

type App struct {
	juejin *JueJin
	mcp    *goMcp.Server
}

func NewApp() *App {
	return &App{
		juejin: NewJueJin(),
		mcp:    mcp.InitMCP(),
	}
}

func Start() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: false,
	})

	app.Use(middleware.Cros())
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":10086"))
}
