package ports

import "ecommerce_order/internal/domain/entity"

type OrderEventPublisher interface {
	Execute(order entity.Order) error
}
