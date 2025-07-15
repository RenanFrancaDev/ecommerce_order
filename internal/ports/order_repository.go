package ports

import "ecommerce_order/internal/entity"

type OrderRepository interface {
	Save(order entity.Order) error
}
