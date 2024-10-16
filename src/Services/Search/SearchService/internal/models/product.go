package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID                 uuid.UUID `json:"id"`
	Title              string    `json:"title"`
	Author             string    `json:"author" `
	Publisher          string    `json:"publisher" `
	PublicationYear    int       `json:"publication_year" `
	Price              float64   `json:"price" `
	Description        string    `json:"description" `
	ImageURL           string    `json:"image_url" `
	SoldQuantity       int       `json:"sold_quantity" `
	AverageRating      float64   `json:"average_rating" `
	DiscountPercentage int       `json:"discount_percentage" `
	ProductType        int       `json:"product_type" `
	IsActive           bool      `json:"is_active" `
	CreatedAt          time.Time `json:"created_at"`
}
