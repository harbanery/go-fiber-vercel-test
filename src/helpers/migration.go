package helpers

import (
	"gofiber-marketplace/src/configs"
	"gofiber-marketplace/src/models"
	"log"
)

func Migration() {
	err := configs.DB.AutoMigrate(
		&models.Product{},
	)

	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}
