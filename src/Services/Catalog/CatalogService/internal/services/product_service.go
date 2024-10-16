package services

import (
	"CatalogService/internal/models"
	"CatalogService/internal/repositories"

	"github.com/google/uuid"
)

type ProductService interface {
	GetProductById(id uuid.UUID) (*models.Product, error)
	GetListProducts(field string, order string, page int, limit int, rating float64, genre_id *uuid.UUID, search_term *string) (*[]models.Product, error)
	CreateProduct(product *models.CreateProductDTO) error
	UpdateProduct(product *models.UpdateProductDTO) error
	DeleteProduct(id uuid.UUID) error
}

type ProductServiceImpl struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{
		productRepo: productRepo,
	}
}

func (p *ProductServiceImpl) GetProductById(id uuid.UUID) (*models.Product, error) {
	return p.productRepo.GetProductById(id)
}

func (p *ProductServiceImpl) CreateProduct(product *models.CreateProductDTO) error {
	return p.productRepo.CreateProduct(product)
}

func (p *ProductServiceImpl) UpdateProduct(product *models.UpdateProductDTO) error {
	return p.productRepo.UpdateProduct(product)
}

func (p *ProductServiceImpl) DeleteProduct(id uuid.UUID) error {
	return p.productRepo.DeleteProduct(id)
}

func (p *ProductServiceImpl) GetListProducts(field string, order string, page int, limit int, rating float64, genre_id *uuid.UUID, search_term *string) (*[]models.Product, error) {
	return p.productRepo.GetListProducts(
		field,
		order,
		page,
		limit,
		rating,
		genre_id,
		search_term,
	)
}
