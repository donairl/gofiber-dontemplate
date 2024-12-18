package models

import (
	"github.com/donairl/gofiber-dontemplate/lib/database"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Weight      float64 `json:"weight"`       // Add weight field
	Unit        string  `json:"unit"`         // Add unit of measurement
	ProductType string  `json:"product_type"` // Add type of product field
}

func GetProductByID(id uint) (Product, error) {
	var product Product
	err := database.Connection.First(&product, id).Error
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func ProductCreate(product *Product) error {
	if result := database.Connection.Create(product); result.Error != nil {
		return result.Error
	}
	return nil
}

func ProductUpdate(product *Product) error {
	if result := database.Connection.Save(product); result.Error != nil {
		return result.Error
	}
	return nil
}

func ProductsFindAll() ([]Product, error) {
	var products []Product
	if err := database.Connection.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func ProductDelete(id uint) error {
	return database.Connection.Delete(&Product{}, id).Error
}
