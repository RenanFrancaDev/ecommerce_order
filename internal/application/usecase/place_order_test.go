// internal/application/usecase/place_order_test.go
package usecase

import (
	"testing"
	"time"

	"ecommerce_order/internal/domain/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrderEventPublisher struct {
	mock.Mock
}

func (m *MockOrderEventPublisher) Execute(order entity.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func TestPlaceOrder_Execute(t *testing.T) {
	t.Run("should place order successfully with complete data", func(t *testing.T) {

		mockPublisher := new(MockOrderEventPublisher)
		placeOrderUC := NewPlaceOrder(mockPublisher)

		order := &entity.Order{
			ClientName:    "Renan",
			ClientEmail:   "renan@example.com",
			ShippingValue: 15.9,
			Address: entity.Address{
				CEP:    12345678,
				Street: "Rua Exemplo",
			},
			PaymentMethod: "CREDIT",
			Items: []entity.Item{
				{
					ItemID:          1,
					ItemDescription: "Camisa Polo 2",
					ItemValue:       59.9,
					ItemQuantity:    2,
					Discount:        10,
				},
			},
		}

		mockPublisher.On("Execute", mock.AnythingOfType("entity.Order")).
			Return(nil).
			Run(func(args mock.Arguments) {
				expected := args.Get(0).(entity.Order)

				assert.NotEmpty(t, expected.OrderID)
				assert.Equal(t, entity.OrderStatusOpen, expected.OrderStatus)
				assert.WithinDuration(t, time.Now(), expected.OrderDate, 5*time.Second)
			})

		placeOrderUC.Execute(order)
	})
}
