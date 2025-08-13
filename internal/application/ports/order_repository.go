package ports

import (
	"context"
	"ecommerce_order/internal/domain/entity"
)

type OrderRepository interface {
	Save(ctx context.Context, order *entity.Order) error
}
