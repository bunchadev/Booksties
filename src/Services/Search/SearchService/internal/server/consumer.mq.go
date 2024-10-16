package server

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	Title string `json:"title"`
}

var test = 1

func (s *Server) setupRabbitMQListener() error {
	channel, err := s.rabbitConn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}

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

	queue, err := channel.QueueDeclare(
		"search_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	err = channel.QueueBind(
		queue.Name,
		"product.*",
		"product_exchange",
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to bind queue: %v", err)
	}

	messages, err := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to consume messages: %v", err)
	}

	go func() {
		for msg := range messages {
			var product Product
			if err := json.Unmarshal(msg.Body, &product); err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}

			// switch msg.RoutingKey {
			// case "product.create":
			// 	log.Printf("Create product: %+v", product)
			// case "product.update":
			// 	log.Printf("Update product: %+v", product)
			// case "product.delete":
			// 	log.Printf("Delete product with ID: %s", product)
			// default:
			// 	log.Printf("Unknown routing key: %s", msg.RoutingKey)
			// }
			if test < 5 {
				log.Printf("return Create product: %+v", product)
				test++
				msg.Nack(false, true)
			} else {
				log.Printf("Create product: %+v", product)
				msg.Ack(false)
			}
		}
	}()

	log.Println("Search service is waiting for messages from Catalog...")
	return nil
}
