package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

type Product struct {
	Title string `json:"title"`
}

func (s *Server) sendProductToQueue(action string, product Product) error {
	channel, err := s.rabbitConn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer channel.Close()
	err = channel.ExchangeDeclare(
		"product_exchange",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %v", err)
	}

	productBytes, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("failed to serialize product: %v", err)
	}

	var routingKey string
	switch action {
	case "create":
		routingKey = "product.create"
	case "update":
		routingKey = "product.update"
	case "delete":
		routingKey = "product.delete"
	default:
		return fmt.Errorf("unknown action: %v", action)
	}

	err = channel.PublishWithContext(
		context.Background(),
		"product_exchange",
		routingKey,
		false,
		false,
		amqp091.Publishing{
			DeliveryMode: amqp091.Persistent,
			ContentType:  "application/json",
			Body:         productBytes,
		})
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}
	return nil
}

func (s *Server) TestMq(c *gin.Context) {
	title := c.Param("title")
	product := Product{
		Title: title,
	}
	if err := s.sendProductToQueue("create", product); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
