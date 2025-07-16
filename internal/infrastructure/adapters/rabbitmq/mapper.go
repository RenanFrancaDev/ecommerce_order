package rabbitmq

import "ecommerce_order/internal/domain/entity"

type OrderMapper struct{}

func NewOrderMapper() *OrderMapper {
	return &OrderMapper{}
}

func (m *OrderMapper) ToDto(order entity.Order) OrderDto {
	return OrderDto{
		OrderID:       order.OrderID,
		OrderDate:     order.OrderDate,
		OrderStatus:   order.OrderStatus,
		ClientName:    order.ClientName,
		ClientEmail:   order.ClientEmail,
		ShippingValue: order.ShippingValue,
		Address:       m.toAddressDto(order.Address),
		PaymentMethod: order.PaymentMethod,
		Items:         m.toItemsDto(order.Items),
	}
}

func (m *OrderMapper) toAddressDto(address entity.Address) AddressDto {
	return AddressDto{
		CEP:    address.CEP,
		Street: address.Street,
	}
}

func (m *OrderMapper) toItemsDto(items []entity.Item) []ItemDto {
	result := make([]ItemDto, len(items))
	for i, item := range items {
		result[i] = ItemDto{
			ItemID:          item.ItemID,
			ItemDescription: item.ItemDescription,
			ItemValue:       item.ItemValue,
			ItemQuantity:    item.ItemQuantity,
			Discount:        item.Discount,
		}
	}
	return result
} 