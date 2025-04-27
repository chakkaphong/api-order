package repo

import (
	"api-order/internal/entities"
	"context"
)

func (r *OrderRepository) UpdateOrderStatusById(ctx context.Context, orderId string, status entities.OrderStatusEnum) error {
	tx := r.postgres.WithContext(ctx)

	if err := tx.
		Model(&entities.Orders{}).
		Where("id = ?", orderId).
		Update("status", status).Error; err != nil {
		return err
	}

	return nil
}
