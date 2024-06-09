package handler

import (
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"required,gt=0"`
	Image       string  `json:"image" validate:"required"`
	Size        uint    `json:"size" validate:"required,gt=0"`
	Color       string  `json:"color" validate:"required,iscolor"`
	Rating      uint    `json:"rating" gorm:"default:0"`
	Description string  `json:"description" validate:"required"`
	CategoryID  uint    `json:"category_id" validate:"required"`
	SellerID    uint    `json:"seller_id" validate:"required"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, This is API for Blanja by Raihan Yusuf from Codecraft")
	})

	app.Get("/products", getAllProducts)

	return adaptor.FiberApp(app)
}

func getAllProducts(c *fiber.Ctx) error {

	url := os.Getenv("URL")
	var err error
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	var products []*Product
	results := db.Find(&products)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":     "success",
		"statusCode": 200,
		"data":       results,
	})
}
