package usecase

import (
	"errors"
	"testing"

	"ecommerce_order/internal/domain/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrderEventPublisher struct {
	mock.Mock
}

func (m *MockOrderEventPublisher) Execute(order *entity.Order) error {
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

		mockPublisher.On("Execute", mock.AnythingOfType("*entity.Order")).
			Return(nil).
			Run(func(args mock.Arguments) {
				expected := args.Get(0).(*entity.Order)

				assert.NotEmpty(t, expected.OrderID)
				assert.Equal(t, entity.OrderStatusOpen, expected.OrderStatus)
				assert.NotEmpty(t, expected.OrderDate)
			})

		err := placeOrderUC.Execute(order)
		assert.Nil(t, err)

	})

	t.Run("should return error when publisher fails", func(t *testing.T) {

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

		expectedError := errors.New("rabbitmq connection failed")
		mockPublisher.On("Execute", mock.AnythingOfType("*entity.Order")).
			Return(expectedError)

		err := placeOrderUC.Execute(order)

		assert.Error(t, err)
		assert.Equal(t, expectedError, err)

		mockPublisher.AssertExpectations(t)
	})

}
