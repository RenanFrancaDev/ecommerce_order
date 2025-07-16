package dto

import "time"

type StandardResponse struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
}

type StandardErrorResponse struct {
	Message   string      `json:"message"`
	Status    string      `json:"status"`
	ErrorCode string      `json:"error_code"`
	Details   interface{} `json:"details,omitempty"`
}

type OrderResponse struct {
	OrderID       string          `json:"order_id"`
	OrderDate     time.Time       `json:"order_date"`
	OrderStatus   string          `json:"order_status"`
	ClientName    string          `json:"client_name"`
	ClientEmail   string          `json:"client_email"`
	ShippingValue float64         `json:"shipping_value"`
	Address       AddressResponse `json:"address"`
	PaymentMethod string          `json:"payment_method"`
	Items         []ItemResponse  `json:"items"`
	TotalValue    float64         `json:"total_value"`
}

type AddressResponse struct {
	CEP    int    `json:"cep"`
	Street string `json:"street"`
}

type ItemResponse struct {
	ItemID          int     `json:"item_id"`
	ItemDescription string  `json:"item_description"`
	ItemValue       float64 `json:"item_value"`
	ItemQuantity    int     `json:"item_quantity"`
	Discount        float64 `json:"discount"`
	TotalValue      float64 `json:"total_value"`
}

type ValidationErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
