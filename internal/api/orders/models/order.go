package models

import (
	"api-order/internal/entities"
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
}

type OrderItem struct {
	ProductName string          `json:"product_name" validate:"required"`
	Quantity    int             `json:"quantity" validate:"required"`
	Price       decimal.Decimal `json:"price" validate:"required"`
}

type CreateOrderRequest struct {
	CustomerName string      `json:"customer_name" validate:"required"`
	OrderItems   []OrderItem `json:"order_items"`
}

type UpdateOrderRequest struct {
	Status entities.OrderStatusEnum `json:"status"`
}

type OrderResponse struct {
	Id           int                      `json:"id"`
	CustomerName string                   `json:"customer_name"`
	TotalAmount  decimal.Decimal          `json:"total_amount"`
	Status       entities.OrderStatusEnum `json:"status"`
	CreatedAt    time.Time                `json:"created_at"`
	OrderItems   []OrderItemsResponse     `json:"order_items"`
}

type OrderItemsResponse struct {
	Id          int             `json:"id"`
	OrderId     int             `json:"order_id"`
	ProductName string          `json:"product_name"`
	Quantity    int             `json:"quantity"`
	Price       decimal.Decimal `json:"price"`
}

type OrderFilter struct {
	Page  int `query:"order_status"`
	Limit int `query:"order_status"`
}
