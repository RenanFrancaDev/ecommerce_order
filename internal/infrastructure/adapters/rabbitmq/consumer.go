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
		false, 
		false, 
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Error to consumer queue: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("[consumer] Shutdown signal received. Exiting...")
			return nil

		case msg, ok := <-msgs:
			if !ok {
				log.Println("[consumer] Message channel closed. Exiting...")
				return nil
			}

			var orderDTO OrderDto
			if err := json.Unmarshal(msg.Body, &orderDTO); err != nil {
				msg.Nack(false, false) // descarta mensagem inválida
				continue
			}

			order := c.mapper.FromDto(orderDTO)

			if _, err := c.mongo.InsertOne(ctx, order); err != nil {
				log.Printf("[mongo] Failed to insert order: %v", err)
				msg.Nack(false, true) // reenvia a mensagem para tentar depois
				continue
			}

			msg.Ack(false)
		}
	}	

}
