package services

import (
	"SearchService/internal/models"
	"SearchService/internal/repositories"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
)

type ProductService interface {
	CreateProduct(product models.Product) (*index.Response, error)
}

type ProductServiceImpl struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{
		productRepo: productRepo,
	}
}

func (p *ProductServiceImpl) CreateProduct(product models.Product) (*index.Response, error) {
	return p.productRepo.CreateProduct(product)
}
