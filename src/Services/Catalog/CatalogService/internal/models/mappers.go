package models

func (p Product) ProductToProductDto() *ProductDto {
	return &ProductDto{
		ID:                 p.ID,
		Title:              p.Title,
		Author:             p.Author,
		Publisher:          p.Publisher,
		PublicationYear:    p.PublicationYear,
		PageCount:          p.PageCount,
		Dimensions:         p.Dimensions,
		CoverType:          p.CoverType,
		Price:              p.Price,
		Description:        p.Description,
		ImageURL:           p.ImageURL,
		SoldQuantity:       p.SoldQuantity,
		AverageRating:      p.AverageRating,
		QuantityEvaluate:   p.QuantityEvaluate,
		DiscountPercentage: p.DiscountPercentage,
	}
}
