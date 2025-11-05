package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/unomcp/JueJin-MCP/middleware"
)

func Start(port string) {
	app := NewApp(fiber.New(fiber.Config{
		DisableStartupMessage: false,
	}))

	app.FiberApp.Use(middleware.Cros())

	SetupRoutes(app)

	log.Fatal(app.FiberApp.Listen(port))
}

func main() {
	Start(":10086")
}
