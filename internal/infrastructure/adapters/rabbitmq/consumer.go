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
		return fmt.Errorf("erro ao abrir canal: %w", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		c.queueName,
		"",
		true,  // Auto-Ack
		false, // exclusive
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("erro ao consumir da fila: %w", err)
	}

	log.Println("🟢 Aguardando mensagens da fila...")

	for msg := range msgs {
		var orderDTO OrderDto
		if err := json.Unmarshal(msg.Body, &orderDTO); err != nil {
			log.Printf("❌ erro ao desserializar: %v", err)
			continue
		}

		order := c.mapper.FromDto(orderDTO)

		if _, err := c.mongo.InsertOne(context.TODO(), order); err != nil {
			log.Printf("❌ erro ao salvar no MongoDB: %v", err)
			continue
		}

		log.Printf("✅ Pedido %s salvo no MongoDB", order.OrderID)
	}

	return nil
}
