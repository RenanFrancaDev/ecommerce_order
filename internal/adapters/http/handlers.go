package http

import (
	"ecommerce_order/internal/config"
)
type Handlers struct {
	Order *OrderHandler
}
func BuildHandlers(cfg *config.Config) *Handlers {
	return &Handlers{
		Order: BuildOrderHandler(cfg),
	}
}