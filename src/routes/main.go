package routes

import (
	"gofiber-marketplace/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, This is API for Blanja by Raihan Yusuf from Codecraft")
	})

	app.Get("/products", controllers.GetAllProduct)
}
