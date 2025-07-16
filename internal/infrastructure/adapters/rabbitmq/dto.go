package rabbitmq

import (
	"time"
)

type AddressDto struct {
	CEP    int    `json:"cep"`
	Street string `json:"street"`
}

type ItemDto struct {
	ItemID          int     `json:"item_id"`
	ItemDescription string  `json:"item_description"`
	ItemValue       float64 `json:"item_value"`
	ItemQuantity    int     `json:"item_quantity"`
	Discount        float64 `json:"discount"`
}

type OrderDto struct {
	OrderID       string      `json:"order_id"`
	OrderDate     time.Time   `json:"order_date"`
	OrderStatus   string      `json:"order_status"`
	ClientName    string      `json:"client_name"`
	ClientEmail   string      `json:"client_email"`
	ShippingValue float64     `json:"shipping_value"`
	Address       AddressDto  `json:"address"`
	PaymentMethod string      `json:"payment_method"`
	Items         []ItemDto   `json:"items"`
} 