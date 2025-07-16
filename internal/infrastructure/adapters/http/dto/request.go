package dto

type CreateOrderRequest struct {
	ClientName    string               `json:"client_name" binding:"required"`
	ClientEmail   string               `json:"client_email" binding:"required,email"`
	ShippingValue float64              `json:"shipping_value" binding:"min=0"`
	Address       CreateAddressRequest `json:"address" binding:"required"`
	PaymentMethod string               `json:"payment_method" binding:"required"`
	Items         []CreateItemRequest  `json:"items" binding:"required,min=1"`
}

type CreateAddressRequest struct {
	CEP    int    `json:"cep" binding:"required,min=1"`
	Street string `json:"street" binding:"required"`
}

type CreateItemRequest struct {
	ItemID          int     `json:"item_id" binding:"required,min=1"`
	ItemDescription string  `json:"item_description" binding:"required"`
	ItemValue       float64 `json:"item_value" binding:"required,min=0.01"`
	ItemQuantity    int     `json:"item_quantity" binding:"required,min=1"`
	Discount        float64 `json:"discount" binding:"min=0"`
}