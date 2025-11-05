package middleware

import "github.com/gofiber/fiber/v2"

func Cros() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept")

		// 处理 OPTIONS 预检请求
		if c.Method() == "OPTIONS" {
			c.Status(204)
			return nil
		}

		return c.Next()
	}
}
