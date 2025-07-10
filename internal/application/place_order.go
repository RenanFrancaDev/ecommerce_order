package application

import (
	"time"

	"ecommerce_order/internal/domain"
	"ecommerce_order/internal/ports"

	"github.com/google/uuid"
)

type OrderService struct {
	Publisher ports.OrderPublisher
}

func NewOrderService(publisher ports.OrderPublisher) *OrderService {
	return &OrderService{Publisher: publisher}
}

func (s *OrderService) PlaceOrder(order *domain.Order) error {
	order.OrderID = uuid.New().String()
	order.OrderDate = time.Now()

	if order.OrderStatus == "" {
		order.OrderStatus = "OPEN"
	}

	return s.Publisher.Publish(*order)
}
