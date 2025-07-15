package entity

import "time"

type Address struct {
	CEP    int    `json:"cep"`
	Street string `json:"street"`
}

type Item struct {
	ItemID          int     `json:"item_id"`
	ItemDescription string  `json:"item_description"`
	ItemValue       float64 `json:"item_value"`
	ItemQuantity    int     `json:"item_quantity"`
	Discount        float64 `json:"discount"`
}

type Order struct {
	OrderID       string    `json:"order_id"`
	OrderDate     time.Time `json:"order_date"`
	OrderStatus   string    `json:"orderStatus"`
	ClientName    string    `json:"client_name"`
	ClientEmail   string    `json:"client_email"`
	ShippingValue float64   `json:"shipping_value"`
	Address       Address   `json:"address"`
	PaymentMethod string    `json:"paymentMethod"`
	Items         []Item    `json:"items"`
}
