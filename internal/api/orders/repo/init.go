package repo

import (
	"api-order/internal/api/orders"
	"api-order/internal/database"
	"context"

	"gorm.io/gorm"
)

type OrderRepository struct {
	postgres *gorm.DB
}

func NewOrderRepository(postgres database.Service) orders.Repository {
	return &OrderRepository{
		postgres: postgres.DB(),
	}
}

func (r *OrderRepository) WithTransaction(ctx context.Context, fn func(rp orders.Repository) error) error {
	// Start a new transaction
	return r.postgres.Transaction(func(tx *gorm.DB) error {
		txRepo := &OrderRepository{postgres: tx}

		if err := fn(txRepo); err != nil {
			return err
		}
		return nil
	})
}
