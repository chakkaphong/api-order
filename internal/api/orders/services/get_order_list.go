package services

import (
	"api-order/internal/entities"
	"api-order/internal/global/responses"
	"context"
)

func (s *OrderServices) GetOrders(ctx context.Context, p responses.Paging) ([]entities.Orders, int64, error) {
	return s.repo.GetOrders(ctx, p)
}
