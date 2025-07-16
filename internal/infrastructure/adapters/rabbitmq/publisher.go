package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"ecommerce_order/internal/application/ports"
	"ecommerce_order/internal/domain/entity"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type Publisher struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	queueName string
	mapper    *OrderMapper
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("[msg:env file not found] [error: %v]", err)
	}
}

func (p *Publisher) Execute(order entity.Order) error {
	orderDto := p.mapper.ToDto(order)

	body, err := json.Marshal(orderDto)
	if err != nil {
		return fmt.Errorf("failed to serialize order: %w", err)
	}

	_, err = p.channel.QueueDeclare(
		p.queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
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
		return fmt.Errorf("failed to publish order to queue: %w", err)
	}

	return nil
}

func NewRabbitMQPublisher(conn *amqp.Connection, queueName string) ports.OrderEventPublisher {
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("[feature:order] [msg:failed to create rabbitmq channel] [error: %v]", err)
		panic("Failed to create RabbitMQ channel")
	}

	return &Publisher{
		conn:      conn,
		channel:   ch,
		queueName: queueName,
		mapper:    NewOrderMapper(),
	}
}
