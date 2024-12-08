package repository

import (
	"github.com/donairl/gofiber-dontemplate/models"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	return models.ProductsFindAll()
}
