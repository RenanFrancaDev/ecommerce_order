package ports

import "ecommerce_order/internal/entity"

type QueuePublisher interface {
	Publish(order entity.Order) error
}


