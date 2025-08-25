package usecase

import (
	"time"

	"ecommerce_order/internal/application/ports"
	"ecommerce_order/internal/domain/entity"

	"github.com/google/uuid"
)

type (
	PlaceOrderUseCase interface {
		Execute(order *entity.Order) error
	}
	PlaceOrder struct {
		publisher ports.OrderEventPublisher
	}
)

func NewPlaceOrder(publisher ports.OrderEventPublisher) PlaceOrderUseCase {
	return &PlaceOrder{
		publisher: publisher,
	}
}

func (uc *PlaceOrder) Execute(order *entity.Order) error {
	order.OrderID = uuid.New().String()
	order.OrderDate = time.Now()
	order.OrderStatus = entity.OrderStatusOpen
	return uc.publisher.Execute(order)
}
