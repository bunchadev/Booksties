package repositories

import (
	"SearchService/internal/models"
	"context"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/index"
)

type ProductRepository interface {
	CreateProduct(product models.Product) (*index.Response, error)
	DeleteProduct(id string) error
}

type ProductRepositoryImpl struct {
	esClient *elasticsearch.TypedClient
}

func NewProductRepository(esClient *elasticsearch.TypedClient) ProductRepository {
	return &ProductRepositoryImpl{
		esClient: esClient,
	}
}

func (p *ProductRepositoryImpl) CreateProduct(product models.Product) (*index.Response, error) {
	product.CreatedAt = time.Now()

	res, err := p.esClient.
		Index("products").
		Id(product.ID.String()).
		Request(product).
		Do(context.TODO())

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *ProductRepositoryImpl) DeleteProduct(id string) error {
	_, err := p.esClient.
		Delete("products", id).
		Do(context.TODO())
	if err != nil {
		return err
	}
	return nil
}
