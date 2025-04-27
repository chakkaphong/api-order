package repo

import (
	"api-order/internal/entities"
	"api-order/internal/global/responses"
	"context"
)

func (r *OrderRepository) GetOrders(ctx context.Context, p responses.Paging) ([]entities.Orders, int64, error) {
	var orders []entities.Orders
	var total int64

	if err := r.postgres.WithContext(ctx).Model(&entities.Orders{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (p.Page - 1) * p.Limit

	if err := r.postgres.WithContext(ctx).
		Limit(p.Limit).
		Offset(offset).
		Preload("OrderItems").Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}
