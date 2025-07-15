package application

import (
	"time"

	"ecommerce_order/internal/entity"
	"ecommerce_order/internal/ports"

	"github.com/google/uuid"
)

type PlaceOrderUseCase struct {
    Publisher ports.QueuePublisher
}

const (
    OrderStatusOpen   = "OPEN"
    OrderStatusPaid   = "PAID"
    OrderStatusClosed = "CLOSED" 
)

func NewPlaceOrderUseCase(publisher ports.QueuePublisher) *PlaceOrderUseCase {
	return &PlaceOrderUseCase{Publisher: publisher}
}

func (uc *PlaceOrderUseCase) PlaceOrder(order *entity.Order) error {
	order.OrderID = uuid.New().String()
	order.OrderDate = time.Now()

	if order.OrderStatus == "" {
		order.OrderStatus = OrderStatusOpen
	}

	return uc.Publisher.Publish(*order)
}
