package repositories

import (
	"CatalogService/internal/models"
	"CatalogService/internal/query"
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const dbTimeout = 3 * time.Second

type ProductRepository interface {
	GetProductById(id uuid.UUID) (*models.Product, error)
	GetListProducts(field string, order string, page int, limit int, rating float64, genre_id *uuid.UUID, search_term *string) (*[]models.Product, error)
	CreateProduct(product *models.CreateProductDTO) error
	UpdateProduct(product *models.UpdateProductDTO) error
	DeleteProduct(id uuid.UUID) error
}

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (p *ProductRepositoryImpl) GetProductById(id uuid.UUID) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var product models.Product
	var imageURL *string
	err := p.db.QueryRowContext(ctx, query.GET_USER_BY_ID, id).Scan(
		&product.ID, &product.Title, &product.Author,
		&product.Publisher, &product.PublicationYear,
		&product.PageCount, &product.Dimensions,
		&product.CoverType, &product.Price, &product.Description,
		&imageURL, &product.SoldQuantity, &product.AverageRating,
		&product.QuantityEvaluate, &product.DiscountPercentage, &product.ProductType,
		&product.IsActive, &product.OriginalOwnerID, &product.CreatedAt,
		&product.UpdatedAt,
	)

	if imageURL != nil {
		product.ImageURL = *imageURL
	} else {
		product.ImageURL = ""
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepositoryImpl) CreateProduct(product *models.CreateProductDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(ctx, query.CREATE_PRODUCT,
		product.Title, product.Author, product.Publisher,
		product.PublicationYear, product.PageCount,
		product.Dimensions, product.CoverType,
		product.Price, product.Description,
		product.ImageURL, product.OriginalOwnerID,
	)

	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (p *ProductRepositoryImpl) UpdateProduct(product *models.UpdateProductDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(ctx, query.UPDATE_PRODUCT,
		product.ID, product.Title, product.Author, product.Publisher,
		product.PublicationYear, product.PageCount, product.Dimensions, product.CoverType,
		product.Price, product.Description, product.DiscountPercentage,
		product.ProductType, product.IsActive,
	)

	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (p *ProductRepositoryImpl) DeleteProduct(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(ctx, query.DELETE_PRODUCT, id)

	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (p *ProductRepositoryImpl) GetListProducts(field string, order string, page int, limit int, rating float64, genre_id *uuid.UUID, search_term *string) (*[]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	rows, err := p.db.QueryContext(ctx, query.PAGINATION_PRODUCT,
		field,
		order,
		page,
		limit,
		rating,
		genre_id,
		search_term,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID, &product.Title, &product.Author,
			&product.Publisher, &product.PublicationYear,
			&product.PageCount, &product.Dimensions,
			&product.CoverType, &product.Price, &product.Description,
			&product.ImageURL, &product.SoldQuantity, &product.AverageRating,
			&product.QuantityEvaluate, &product.DiscountPercentage, &product.ProductType,
			&product.IsActive, &product.OriginalOwnerID, &product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return &products, nil
}
