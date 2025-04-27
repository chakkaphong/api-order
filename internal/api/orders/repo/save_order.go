package repo

import (
	"api-order/internal/entities"
	"context"
	"log"
)

func (r *OrderRepository) SaveOrder(ctx context.Context, order entities.Orders) (*entities.Orders, error) {
	if ctx.Err() != nil {
		log.Printf("Context already canceled before sending to worker: %v", ctx.Err())
	}

	tx := r.postgres.WithContext(ctx)

	if err := tx.Create(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
