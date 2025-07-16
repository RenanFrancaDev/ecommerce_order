package http

import (
	"ecommerce_order/internal/infrastructure/container"
)

type Handlers struct {
	Order *OrderHandler
}

func NewHandlers(container *container.Container) *Handlers {
	return &Handlers{
		Order: NewOrderHandler(container.GetPlaceOrderUseCase()),
	}
}
