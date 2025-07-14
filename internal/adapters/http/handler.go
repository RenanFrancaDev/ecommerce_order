package http

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"

	"ecommerce_order/internal/adapters/rabbitmq"
	"ecommerce_order/internal/application"
	"ecommerce_order/internal/domain"
)

func RegisterRoutes(r *gin.Engine) {

	amqpURL := os.Getenv("RABBITMQ_URL")
	if amqpURL == "" {
		log.Fatal("RABBITMQ_URL not configured in .env file")
	}


	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatal("Eror to connect RabbitMQ:", err)
	}

	publisher := rabbitmq.NewRabbitMQPublisher(conn, "orders")

	orderService := application.NewPlaceOrderUseCase(publisher)

	r.POST("/orders", func(c *gin.Context) {
		var order domain.Order

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
			return
		}

		if err := orderService.PlaceOrder(&order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Order processing error", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Order processed successfully",
			"order_id": order.OrderID,
		})
	})
}