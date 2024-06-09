package controllers

import (
	"gofiber-marketplace/src/helpers"
	"gofiber-marketplace/src/models"
	"math"

	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(c *fiber.Ctx) error {
	keyword := c.Query("search")
	sort := helpers.GetSortParams(c.Query("sorting"), c.Query("orderBy"))
	page, limit, offset := helpers.GetPaginationParams(c.Query("limit"), c.Query("page"))
	totalData := models.CountData(keyword)
	totalPage := math.Ceil(float64(totalData) / float64(limit))

	products := models.SelectAllProducts(keyword, sort, limit, offset)
	if len(products) == 0 {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status":     "no content",
			"statusCode": 202,
			"message":    "Product is empty. You should create product",
		})
	}

	resultProducts := make([]*map[string]interface{}, len(products))
	for i, product := range products {
		resultProducts[i] = &map[string]interface{}{
			"id":          product.ID,
			"created_at":  product.CreatedAt,
			"updated_at":  product.UpdatedAt,
			"category_id": product.CategoryID,
			"brand_id":    product.SellerID,
			"name":        product.Name,
			"photo":       product.Image,
			"rating":      product.Rating,
			"price":       product.Price,
			"size":        product.Size,
			"color":       product.Color,
			"stock":       product.Stock,
			"condition":   product.Condition,
			"desc":        product.Description,
		}
	}

	// return c.JSON(products)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "success",
		"statusCode":  200,
		"data":        resultProducts,
		"currentPage": page,
		"limit":       limit,
		"totalData":   totalData,
		"totalPage":   totalPage,
	})
}
