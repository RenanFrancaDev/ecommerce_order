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

		mockPublisher.On("Execute", mock.AnythingOfType("entity.Order")).
			Return(nil).
			Run(func(args mock.Arguments) {
				orderArg := args.Get(0).(entity.Order)

				expected := *order
				expected.OrderStatus = entity.OrderStatusOpen //POR QUE esse campo não pode ser declarado no order?

				assertOrderEqual(t, expected, orderArg, true)
			})

		// Act
		placeOrderUC.Execute(order)
	})
}

// helper function
func assertOrderEqual(t *testing.T, expected, actual entity.Order, checkDynamicFields bool) {
	if checkDynamicFields {
		assert.NotEmpty(t, actual.OrderID)
		assert.Equal(t, entity.OrderStatusOpen, actual.OrderStatus)
		assert.WithinDuration(t, time.Now(), actual.OrderDate, time.Second)
	}

	assert.Equal(t, expected.ClientName, actual.ClientName)
	assert.Equal(t, expected.ClientEmail, actual.ClientEmail)
	assert.Equal(t, expected.ShippingValue, actual.ShippingValue)
	assert.Equal(t, expected.Address.CEP, actual.Address.CEP)
	assert.Equal(t, expected.Address.Street, actual.Address.Street)
	assert.Equal(t, expected.PaymentMethod, actual.PaymentMethod)
	assert.Equal(t, expected.TotalValue(), actual.TotalValue())

	assert.Len(t, actual.Items, len(expected.Items))
	for i := range expected.Items {
		assert.Equal(t, expected.Items[i].ItemID, actual.Items[i].ItemID)
		assert.Equal(t, expected.Items[i].ItemDescription, actual.Items[i].ItemDescription)
		assert.Equal(t, expected.Items[i].ItemValue, actual.Items[i].ItemValue)
		assert.Equal(t, expected.Items[i].ItemQuantity, actual.Items[i].ItemQuantity)
		assert.Equal(t, expected.Items[i].Discount, actual.Items[i].Discount)
		assert.Equal(t, expected.Items[i].TotalValue(), actual.Items[i].TotalValue())
	}
}
