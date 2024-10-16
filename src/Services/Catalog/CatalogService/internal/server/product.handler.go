package server

import (
	"CatalogService/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) GetProductByIDHander(c *gin.Context) {
	productID := c.Param("id")
	id, err := uuid.Parse(productID)
	if err != nil {
		models.ErrorResponse(c, 301, "Parameter is invalid")
		return
	}
	product, err := s.productService.GetProductById(id)
	if err != nil {
		models.ErrorResponse(c, 301, err.Error())
		return
	}
	models.SuccessResponse(c, 201, "Get product success", product.ProductToProductDto())
}

func (s *Server) CreateProductHandler(c *gin.Context) {
	var dto models.CreateProductDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		models.ErrorResponse(c, 301, err.Error())
		return
	}
	err := s.productService.CreateProduct(&dto)
	if err != nil {
		models.ErrorResponse(c, 301, err.Error())
		return
	}
	models.SuccessResponse(c, 201, "Create success", dto.Title)
}

func (s *Server) UpdateProductHandler(c *gin.Context) {
	var dto models.UpdateProductDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		models.ErrorResponse(c, 301, err.Error())
		return
	}
	err := s.productService.UpdateProduct(&dto)
	if err != nil {
		models.ErrorResponse(c, 301, err.Error())
		return
	}
	models.SuccessResponse(c, 201, "Update success", dto.Title)
}

func (s *Server) DeleteProductHandler(c *gin.Context) {
	productID := c.Param("id")
	id, err := uuid.Parse(productID)
	if err != nil {
		models.ErrorResponse(c, 301, "Parameter is invalid")
		return
	}
	err = s.productService.DeleteProduct(id)
	if err != nil {
		models.ErrorResponse(c, 301, err.Error())
		return
	}
	models.SuccessResponse(c, 201, "Delete product success", id)
}

func (s *Server) GetListProductsHandler(c *gin.Context) {
	field := c.DefaultQuery("field", "id")
	order := c.DefaultQuery("order", "asc")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		models.ErrorResponse(c, 301, "Invalid page")
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "8"))
	if err != nil {
		models.ErrorResponse(c, 301, "Invalid limit")
		return
	}

	rating, err := strconv.ParseFloat(c.DefaultQuery("rating", "0"), 64)
	if err != nil {
		models.ErrorResponse(c, 301, "Invalid rating")
		return
	}

	genreID := c.Query("genre_id")
	var genreUUID *uuid.UUID
	if genreID != "" {
		parsedGenreID, err := uuid.Parse(genreID)
		if err != nil {
			models.ErrorResponse(c, 301, "Invalid genre ID")
			return
		}
		genreUUID = &parsedGenreID
	} else {
		genreUUID = nil
	}

	searchTerm := c.Query("search_term")
	var searchString *string
	if searchTerm != "" {
		searchString = &searchTerm
	} else {
		searchString = nil
	}
	products, err := s.productService.GetListProducts(
		field,
		order,
		page,
		limit,
		rating,
		genreUUID,
		searchString,
	)
	if err != nil {
		models.ErrorResponse(c, 301, err.Error())
		return
	}
	models.PaginateResponse(c, 201, "Get list products success", page, limit, len(*products), products)
}
