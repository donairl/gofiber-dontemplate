package services

import (
	"github.com/donairl/gofiber-dontemplate/models"
	"github.com/donairl/gofiber-dontemplate/repository"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{
		repo: repo,
	}
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.FindAll()
}
