// internal/application/usecase/place_order_test.go
package usecase

import (
	"testing"
	"time"

	"ecommerce_order/internal/domain/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockOrderEventPublisher to test
type MockOrderEventPublisher struct {
	mock.Mock
}

func (m *MockOrderEventPublisher) Execute(order entity.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func TestPlaceOrder_Execute(t *testing.T) {
	t.Run("should place order successfully with complete data", func(t *testing.T) {
		
		// Arrange
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

		// Calculate expected total
		expectedTotal := order.TotalValue()

		mockPublisher.On("Execute", mock.AnythingOfType("entity.Order")).
			Return(nil).
			Run(func(args mock.Arguments) {
				orderArg := args.Get(0).(entity.Order)
				
				// Verify that order was properly populated
				assert.NotEmpty(t, orderArg.OrderID)
				assert.Equal(t, entity.OrderStatusOpen, orderArg.OrderStatus)
				assert.WithinDuration(t, time.Now(), orderArg.OrderDate, time.Second)
				
				// Verify original data is preserved
				assert.Equal(t, "Renan", orderArg.ClientName)
				assert.Equal(t, "renan@example.com", orderArg.ClientEmail)
				assert.Equal(t, 15.9, orderArg.ShippingValue)
				assert.Equal(t, 12345678, orderArg.Address.CEP)
				assert.Equal(t, "Rua Exemplo", orderArg.Address.Street)
				assert.Equal(t, "CREDIT", orderArg.PaymentMethod)
				assert.Len(t, orderArg.Items, 1)
				
				// Verify calculated values
				assert.Equal(t, expectedTotal, orderArg.TotalValue())
				assert.Equal(t, 109.8, orderArg.Items[0].TotalValue()) // (59.9 * 2) - 10
			})

		// Act
		err := placeOrderUC.Execute(order)

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, order.OrderID)
		assert.Equal(t, entity.OrderStatusOpen, order.OrderStatus)
		assert.WithinDuration(t, time.Now(), order.OrderDate, time.Second)
		assert.Equal(t, expectedTotal, order.TotalValue())
		
		mockPublisher.AssertExpectations(t)
	})

	
}