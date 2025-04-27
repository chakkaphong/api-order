package repo

import (
	"api-order/internal/entities"
	"context"
)

func (r *OrderRepository) GetOrder(ctx context.Context, orderId string) (*entities.Orders, error) {
	var order entities.Orders

	if err := r.postgres.WithContext(ctx).
		Where("id = ?", orderId).
		Preload("OrderItems").
		First(&order).
		Error; err != nil {
		return nil, err
	}

	return &order, nil
}
