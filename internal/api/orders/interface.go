package orders

import (
	"api-order/internal/api/orders/models"
	"api-order/internal/entities"
	"api-order/internal/global/responses"
	"context"
)

type Service interface {
	CreateOrder(ctx context.Context, req models.CreateOrderRequest) error
	GetOrders(ctx context.Context, p responses.Paging) ([]entities.Orders, int64, error)
	GetOrderById(ctx context.Context, orderId string) (*entities.Orders, error)
	UpdateOrder(ctx context.Context, orderId string, req models.UpdateOrderRequest) (*entities.Orders, error)
}

type Repository interface {
	WithTransaction(ctx context.Context, fn func(rp Repository) error) error

	// Order
	SaveOrder(ctx context.Context, order entities.Orders) (*entities.Orders, error)
	UpdateOrderStatusById(ctx context.Context, orderId string, status entities.OrderStatusEnum) error
	GetOrder(ctx context.Context, orderId string) (*entities.Orders, error)
	GetOrders(ctx context.Context, p responses.Paging) ([]entities.Orders, int64, error)

	// OrderItem
	SaveOrderItem(ctx context.Context, orderItems []entities.OrderItems) ([]entities.OrderItems, error)
}
