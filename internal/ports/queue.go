package ports

import "ecommerce_order/internal/domain"

type QueuePublisher interface {
	Publish(order domain.Order) error
}


