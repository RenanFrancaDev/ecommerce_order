package mapper

import (
	"ecommerce_order/internal/domain/entity"
	"ecommerce_order/internal/infrastructure/adapters/http/dto"
)

type OrderMapper struct{}

func NewOrderMapper() *OrderMapper {
	return &OrderMapper{}
}

func (m *OrderMapper) ToEntity(req dto.CreateOrderRequest) entity.Order {
	return entity.Order{
		ClientName:    req.ClientName,
		ClientEmail:   req.ClientEmail,
		ShippingValue: req.ShippingValue,
		Address:       m.toAddressEntity(req.Address),
		PaymentMethod: req.PaymentMethod,
		Items:         m.toItemsEntity(req.Items),
	}
}

func (m *OrderMapper) ToResponse(order entity.Order) dto.OrderResponse {
	return dto.OrderResponse{
		OrderID:       order.OrderID,
		OrderDate:     order.OrderDate,
		OrderStatus:   order.OrderStatus,
		ClientName:    order.ClientName,
		ClientEmail:   order.ClientEmail,
		ShippingValue: order.ShippingValue,
		Address:       m.toAddressResponse(order.Address),
		PaymentMethod: order.PaymentMethod,
		Items:         m.toItemsResponse(order.Items),
		TotalValue:    order.TotalValue(),
	}
}

func (m *OrderMapper) ToCreateOrderResponse(order entity.Order) dto.StandardResponse {
	return dto.StandardResponse{
		Message: "Order processed successfully",
		Status:  "success",
		Data:    m.ToResponse(order),
	}
}

func (m *OrderMapper) toAddressEntity(req dto.CreateAddressRequest) entity.Address {
	return entity.Address{
		CEP:    req.CEP,
		Street: req.Street,
	}
}

func (m *OrderMapper) toAddressResponse(addr entity.Address) dto.AddressResponse {
	return dto.AddressResponse{
		CEP:    addr.CEP,
		Street: addr.Street,
	}
}

func (m *OrderMapper) toItemsEntity(items []dto.CreateItemRequest) []entity.Item {
	result := make([]entity.Item, len(items))
	for i, item := range items {
		result[i] = entity.Item{
			ItemID:          item.ItemID,
			ItemDescription: item.ItemDescription,
			ItemValue:       item.ItemValue,
			ItemQuantity:    item.ItemQuantity,
			Discount:        item.Discount,
		}
	}
	return result
}

func (m *OrderMapper) toItemsResponse(items []entity.Item) []dto.ItemResponse {
	result := make([]dto.ItemResponse, len(items))
	for i, item := range items {
		result[i] = dto.ItemResponse{
			ItemID:          item.ItemID,
			ItemDescription: item.ItemDescription,
			ItemValue:       item.ItemValue,
			ItemQuantity:    item.ItemQuantity,
			Discount:        item.Discount,
			TotalValue:      item.TotalValue(),
		}
	}
	return result
}
