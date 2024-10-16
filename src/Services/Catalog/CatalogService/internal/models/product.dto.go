package models

import "github.com/google/uuid"

type ProductDto struct {
	ID                 uuid.UUID `json:"id"`
	Title              string    `json:"title"`
	Author             string    `json:"author"`
	Publisher          string    `json:"publisher"`
	PublicationYear    int       `json:"publication_year"`
	PageCount          int       `json:"page_count"`
	Dimensions         string    `json:"dimensions"`
	CoverType          string    `json:"cover_type"`
	Price              float64   `json:"price"`
	Description        string    `json:"description"`
	ImageURL           string    `json:"image_url"`
	SoldQuantity       int       `json:"sold_quantity"`
	AverageRating      float64   `json:"average_rating"`
	QuantityEvaluate   int       `json:"quantity_evaluate"`
	DiscountPercentage int       `json:"discount_percentage"`
}

type CreateProductDTO struct {
	Title           string     `json:"title" binding:"required"`
	Author          string     `json:"author"`
	Publisher       string     `json:"publisher"`
	PublicationYear int        `json:"publication_year"`
	PageCount       int        `json:"page_count"`
	Dimensions      string     `json:"dimensions"`
	CoverType       string     `json:"cover_type"`
	Price           float64    `json:"price"`
	Description     string     `json:"description"`
	ImageURL        *string    `json:"image_url"`
	OriginalOwnerID *uuid.UUID `json:"original_owner_id"`
}

type UpdateProductDTO struct {
	ID                 uuid.UUID `json:"id"`
	Title              string    `json:"title"`
	Author             string    `json:"author"`
	Publisher          string    `json:"publisher"`
	PublicationYear    int       `json:"publication_year"`
	PageCount          int       `json:"page_count"`
	Dimensions         string    `json:"dimensions"`
	CoverType          string    `json:"cover_type"`
	Price              float64   `json:"price"`
	Description        string    `json:"description"`
	DiscountPercentage int       `json:"discount_percentage"`
	ProductType        int       `json:"product_type"`
	IsActive           bool      `json:"is_active"`
}
