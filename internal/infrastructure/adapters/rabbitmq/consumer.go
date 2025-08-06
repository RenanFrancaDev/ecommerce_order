package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
	mongo     *mongo.Collection
	mapper    *OrderMapper
}

func NewConsumer(conn *amqp.Connection, queueName string, collection *mongo.Collection) *Consumer {
	return &Consumer{
		conn:      conn,
		queueName: queueName,
		mongo:     collection,
		mapper:    NewOrderMapper(),
	}
}

func (c *Consumer) Consume(ctx context.Context) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return fmt.Errorf("Error to open channel: %w", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		c.queueName,
		"",
		true, 
		false, 
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Error to consumer queue: %w", err)
	}

	log.Println("🟢 Waiting for queue messages...")

	for msg := range msgs {
		var orderDTO OrderDto
		if err := json.Unmarshal(msg.Body, &orderDTO); err != nil {
			log.Printf("❌ Error while deserializing: %v", err)
			continue
		}

		order := c.mapper.FromDto(orderDTO)

		if _, err := c.mongo.InsertOne(ctx, order); err != nil {
			log.Printf("❌ Error saving to MongoDB: %v", err)
			continue
		}

		log.Printf("✅ Order %s saved to MongoDB", order.OrderID)
	}

	return nil
}
