package juejin

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/unomcp/JueJin-MCP/middleware"
	"github.com/unomcp/JueJin-MCP/routes"
)

func Start() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: false,
	})

	// 添加请求日志中间件
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("[%s] %s", c.Method(), c.Path())
		return c.Next()
	})

	app.Use(middleware.Cros())
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":10086"))
}
