package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rabbitmq/amqp091-go"

	"CatalogService/internal/database"
	"CatalogService/internal/repositories"
	"CatalogService/internal/services"
)

type Server struct {
	port           int
	db             database.Service
	productService services.ProductService
	rabbitConn     *amqp091.Connection
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := database.New()

	rabbitConn, err := amqp091.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		return nil
	}

	NewServer := &Server{
		port:       port,
		db:         db,
		rabbitConn: rabbitConn,
		productService: services.NewProductService(
			repositories.NewProductRepository(db.DB()),
		),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
