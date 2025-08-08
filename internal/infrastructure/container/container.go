package container

import (
	"context"
	"log"

	"ecommerce_order/internal/application/ports"
	"ecommerce_order/internal/application/usecase"
	"ecommerce_order/internal/infrastructure/adapters/rabbitmq"
	"ecommerce_order/internal/infrastructure/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/streadway/amqp"
)

type Container struct {
	cfg               *config.Config
	orderPublisher    ports.OrderEventPublisher
	placeOrderUseCase usecase.PlaceOrderUseCase
	orderConsumer 	  *rabbitmq.Consumer
	mongoClient       *mongo.Client 
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

func (c *Container) GetMongoClient() *mongo.Client {
	if c.mongoClient == nil {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.cfg.MongoURI))
		if err != nil {
			log.Fatalf("Error connecting to MongoDB: %v", err)
		}
		c.mongoClient = client
	}
	return c.mongoClient
}


func (c *Container) GetOrderConsumer() *rabbitmq.Consumer {
	if c.orderConsumer == nil {
		conn, err := amqp.Dial(c.cfg.RabbitMQURL)
		if err != nil {
			log.Fatalf("Error connecting to RabbitMQ: %v", err)
		}

		mongoClient := c.GetMongoClient()

		collection := mongoClient.Database(c.cfg.MongoDatabase).Collection("orders") 
		
		c.orderConsumer = rabbitmq.NewConsumer(conn, c.cfg.RabbitMQOrdersQueue, collection)
	}
	return c.orderConsumer
}

