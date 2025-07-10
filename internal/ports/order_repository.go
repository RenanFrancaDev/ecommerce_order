package ports

import "ecommerce_order/internal/domain"

type OrderRepository interface {
	Save(order domain.Order) error
}
