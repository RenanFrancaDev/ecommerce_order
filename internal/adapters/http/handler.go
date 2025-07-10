package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ecommerce_order/internal/adapters/rabbitmq"
	"ecommerce_order/internal/application"
	"ecommerce_order/internal/domain"
)

func RegisterRoutes(r *gin.Engine) {
	orderService := application.NewOrderService(rabbitmq.Publisher{}) // injeção da implementação

	r.POST("/orders", func(c *gin.Context) {
		var order domain.Order

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
			return
		}

		if err := orderService.PlaceOrder(&order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar ordem"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Ordem enviada com sucesso", "order_id": order.OrderID})
	})
}
