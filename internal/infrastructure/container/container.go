package container

import (
	"log"

	"ecommerce_order/internal/application/ports"
	"ecommerce_order/internal/application/usecase"
	"ecommerce_order/internal/infrastructure/adapters/rabbitmq"
	"ecommerce_order/internal/infrastructure/config"

	"github.com/streadway/amqp"
)

type Container struct {
	cfg               *config.Config
	orderPublisher    ports.OrderEventPublisher
	placeOrderUseCase usecase.PlaceOrderUseCase
}

func NewContainer(cfg *config.Config) *Container {
	return &Container{
		cfg: cfg,
	}
}

func (c *Container) GetOrderPublisher() ports.OrderEventPublisher {
	if c.orderPublisher == nil {
		conn, err := amqp.Dial(c.cfg.RabbitMQURL)
		if err != nil {
			log.Printf("[feature:order] [msg:failed to connect to rabbitmq] [error: %v]", err)
			panic("Failed to connect to RabbitMQ")
		}

		if c.cfg.RabbitMQOrdersQueue == "" {
			log.Printf("[feature:order] [msg:orders queue environment variable not set] [error: %v]", "ORDERS_QUEUE is empty")
			panic("ORDERS_QUEUE environment variable is not set")
		}

		c.orderPublisher = rabbitmq.NewRabbitMQPublisher(conn, c.cfg.RabbitMQOrdersQueue)
	}
	return c.orderPublisher
}

func (c *Container) GetPlaceOrderUseCase() usecase.PlaceOrderUseCase {
	if c.placeOrderUseCase == nil {
		c.placeOrderUseCase = usecase.NewPlaceOrder(c.GetOrderPublisher())
	}
	return c.placeOrderUseCase
}
