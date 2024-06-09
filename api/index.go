package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	app := fiber.New()

	// configs.InitDB()
	// helpers.Migration()
	// routes.Router(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, This is API for Blanja by Raihan Yusuf from Codecraft")
	})

	return adaptor.FiberApp(app)
}
