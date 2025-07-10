package ports

import "ecommerce_order/internal/domain"

type OrderPublisher interface {
	Publish(order domain.Order) error
}
