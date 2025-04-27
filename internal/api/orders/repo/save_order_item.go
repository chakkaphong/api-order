package repo

import (
	"api-order/internal/entities"
	"context"
	"log"
)

func (r *OrderRepository) SaveOrderItem(ctx context.Context, orderItems []entities.OrderItems) ([]entities.OrderItems, error) {
	if ctx.Err() != nil {
		log.Printf("Context already canceled before sending to worker: %v", ctx.Err())
	}

	tx := r.postgres.WithContext(ctx)

	if err := tx.Create(&orderItems).Error; err != nil {
		return nil, err
	}

	return orderItems, nil
}
