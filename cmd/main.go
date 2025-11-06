package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/unomcp/JueJin-MCP/app"
)

func start(port string) {
	app := app.NewApp(fiber.New(fiber.Config{
		DisableStartupMessage: false,
	}))

	app.SetupRoutes()

	log.Fatal(app.FiberApp.Listen(port))
}

func main() {
	start(":10086")
}
