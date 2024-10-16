package server

import (
	"SearchService/internal/repositories"
	"SearchService/internal/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rabbitmq/amqp091-go"
)

type Server struct {
	port           int
	rabbitConn     *amqp091.Connection
	productService services.ProductService
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	rabbitConn, err := amqp091.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return nil
	}

	typedClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTICSEARCH_URL"),
		},
	})

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	NewServer := &Server{
		port:       port,
		rabbitConn: rabbitConn,
		productService: services.NewProductService(
			repositories.NewProductRepository(typedClient),
		),
	}

	err = NewServer.setupRabbitMQListener()
	if err != nil {
		log.Fatalf("Failed to setup RabbitMQ listener: %v", err)
		return nil
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
