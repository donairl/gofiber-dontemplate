package models

import (
	"time"

	"github.com/donairl/gofiber-dontemplate/lib/database"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Price       float64   `json:"price" binding:"required,min=0"`
	ExpiryDate  time.Time `json:"expiry_date"`  // Add expired date field
	Weight      float64   `json:"weight"`       // Add weight field
	ProductType string    `json:"product_type"` // Add type of product field
}

func GetProductByID(id uint) (Product, error) {
	var product Product
	err := database.DB.First(&product, id).Error
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func CreateProduct(product *Product) error {
	if result := database.Connection.Create(product); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllProducts() ([]Product, error) {
	var products []Product
	if err := database.Connection.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
