package server

import (
	"CatalogService/internal/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Static("/images", "./images")

	r.GET("/rabbit-mq/:title", s.TestMq)
	productRoutes := r.Group("api/v1/Product")
	productRoutes.Use(middlewares.JWTAuthMiddleware())
	{
		productRoutes.GET(
			"/:id",
			middlewares.RequirePermission("view_product"),
			s.GetProductByIDHander,
		)
		productRoutes.POST("/save", middlewares.RequirePermission("create_product"), s.CreateProductHandler)
		productRoutes.POST("/update", middlewares.RequirePermission("update_product"), s.UpdateProductHandler)
		productRoutes.GET("/delete/:id", middlewares.RequirePermission("delete_product"), s.DeleteProductHandler)
		productRoutes.GET("/paginate", middlewares.RequirePermission("view_products"), s.GetListProductsHandler)
	}

	return r
}
