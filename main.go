package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/streadway/amqp" //pacote do RabbitMQ
)

// Struct que representa uma ordem de compra
type Address struct {
	CEP    int    `json:"cep"`
	Street string `json:"street"`
}

type Item struct {
	ItemID          int     `json:"item_id"`
	ItemDescription string  `json:"item_description"`
	ItemValue       float64 `json:"item_value"`
	ItemQuantity    int     `json:"item_quantity"`
	Discount        float64 `json:"discount"`
}

type Order struct {
	OrderID       string    `json:"order_id"` // gerado no back
	OrderDate     time.Time `json:"order_date"`
	OrderStatus   string    `json:"orderStatus"` // OPEN, PAYED, FINISH
	ClientName    string    `json:"client_name"`
	ClientEmail   string    `json:"client_email"`
	ShippingValue float64   `json:"shipping_value"`
	Address       Address   `json:"address"`
	PaymentMethod string    `json:"paymentMethod"` // CREDIT, DEBIT, CASH
	Items         []Item    `json:"items"`
}

func main() {
	r := gin.Default()

	// Rota POST /order
	r.POST("/orders", func(c *gin.Context) {
		var order Order
	
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
			return
		}
	
		// Gera UUID para order_id
		order.OrderID = uuid.New().String()
	
		// Define data atual
		order.OrderDate = time.Now()
	
		// Define status default se não vier no JSON
		if order.OrderStatus == "" {
			order.OrderStatus = "OPEN"
		}
	
		// Conexão com RabbitMQ
		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			log.Println("Erro ao conectar no RabbitMQ:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao conectar na fila"})
			return
		}
		defer conn.Close()
	
		ch, err := conn.Channel()
		if err != nil {
			log.Println("Erro ao abrir canal:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro no canal"})
			return
		}
		defer ch.Close()
	
		q, err := ch.QueueDeclare("orders", false, false, false, false, nil)
		if err != nil {
			log.Println("Erro ao declarar fila:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao declarar fila"})
			return
		}
	
		body, err := json.Marshal(order)
		if err != nil {
			log.Println("Erro ao serializar ordem:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao preparar mensagem"})
			return
		}
	
		err = ch.Publish("", q.Name, false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
		if err != nil {
			log.Println("Erro ao publicar:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar mensagem"})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{"status": "Ordem enviada com sucesso", "order_id": order.OrderID})
	})
	

	r.Run(":8080")
}
