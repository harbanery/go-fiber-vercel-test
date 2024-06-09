package routes

import (
	"gofiber-marketplace/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/products", controllers.GetAllProduct)
}
