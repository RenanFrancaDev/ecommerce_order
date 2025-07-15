package http

import (
	"ecommerce_order/internal/adapters/rabbitmq"
	"ecommerce_order/internal/config"
	"ecommerce_order/internal/entity"
	"ecommerce_order/internal/ports"
	application "ecommerce_order/internal/usecase"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type OrderHandler struct {
	usecase *application.PlaceOrderUseCase
}

func BuildOrderHandler(cfg *config.Config) *OrderHandler {
	conn, err := amqp.Dial(cfg.RabbitMQURL)
	if err != nil {
		panic("RabbitMQ erro: " + err.Error())
	}
	var publisher ports.QueuePublisher = rabbitmq.NewRabbitMQPublisher(conn, os.Getenv("QUEUE_NAME"))
	usecase := application.NewPlaceOrderUseCase(publisher)
	return &OrderHandler{usecase: usecase}

}

func (h *OrderHandler) PlaceOrder(ctx *gin.Context) {
	var order entity.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data", "details": err.Error()})
		return
	}
	if err := h.usecase.PlaceOrder(&order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Order processing error", "details": err.Error()}) 
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
			"message": "Order processed successfully",
			"order_id": order.OrderID,
		}) 
}




