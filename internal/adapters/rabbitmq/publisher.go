package rabbitmq

import (
	"encoding/json"
	"log"

	"ecommerce_order/internal/domain"
	"ecommerce_order/internal/ports"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	queueName string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func (p *RabbitMQPublisher) Publish(order domain.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		log.Printf("[RabbitMQ] Failed to serialize order %s: %v", order.OrderID, err)
		return err
	}

	err = p.channel.Publish(
		"",
		p.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Printf("[RabbitMQ] Failed to publish order %s: %v", order.OrderID, err)
		return err
	}

	log.Printf("[RabbitMQ] Order %s successfully published to queue %s", order.OrderID, p.queueName)
	return nil
}

func NewRabbitMQPublisher(conn *amqp.Connection, queueName string) ports.QueuePublisher {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("[RabbitMQ] Failed to create channel:", err)
	}

	return &RabbitMQPublisher{
		conn:      conn,
		channel:   ch,
		queueName: queueName,
	}
}