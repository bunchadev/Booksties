package server

import (
	"SearchService/internal/models"

	"github.com/gin-gonic/gin"
)

func (s *Server) CreateProductHandler(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		models.ErrorResponse(c, 301, err.Error())
		return
	}
	res, err := s.productService.CreateProduct(product)
	if err != nil {
		models.ErrorResponse(c, 301, err.Error())
		return
	}
	models.SuccessResponse(c, 301, "Create product success", res)
}
