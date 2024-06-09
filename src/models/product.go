package models

import (
	"gofiber-marketplace/src/configs"

	"gorm.io/gorm"
)

type ProductCondition string

const (
	New  ProductCondition = "new"
	Used ProductCondition = "used"
)

type Product struct {
	gorm.Model
	Name        string           `json:"name" validate:"required"`
	Price       float64          `json:"price" validate:"required,gt=0"`
	Stock       int              `json:"stock" validate:"required,gt=0"`
	Image       string           `json:"image" validate:"required"`
	Size        uint             `json:"size" validate:"required,gt=0"`
	Color       string           `json:"color" validate:"required,iscolor"`
	Rating      uint             `json:"rating" gorm:"default:0"`
	Description string           `json:"description" validate:"required"`
	Condition   ProductCondition `gorm:"type:product_condition;default:new" json:"condition" validate:"oneof=new used"`
	CategoryID  uint             `json:"category_id" validate:"required"`
	SellerID    uint             `json:"seller_id" validate:"required"`
}

func SelectAllProducts(keyword, sort string, limit, offset int) []*Product {
	var products []*Product
	keyword = "%" + keyword + "%"
	configs.DB.Order(sort).Limit(limit).Offset(offset).Where("name ILIKE ?", keyword).Find(&products)
	return products
}

func CountData(keyword string) int64 {
	var result int64
	keyword = "%" + keyword + "%"
	configs.DB.Table("products").Where("deleted_at IS NULL AND name ILIKE ?", keyword).Count(&result)
	return result
}
