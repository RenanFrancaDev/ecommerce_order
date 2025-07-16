package entity

import (
	"time"
)

const (
	OrderStatusOpen = "OPEN"
)

type Address struct {
	CEP    int
	Street string
}

type Item struct {
	ItemID          int
	ItemDescription string
	ItemValue       float64
	ItemQuantity    int
	Discount        float64
}

func (i Item) TotalValue() float64 {
	return (i.ItemValue * float64(i.ItemQuantity)) - i.Discount
}

type Order struct {
	OrderID       string
	OrderDate     time.Time
	OrderStatus   string
	ClientName    string
	ClientEmail   string
	ShippingValue float64
	Address       Address
	PaymentMethod string
	Items         []Item
}

func (o *Order) TotalValue() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += item.TotalValue()
	}
	return total + o.ShippingValue
}
