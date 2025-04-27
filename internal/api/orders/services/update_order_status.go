package services

import (
	"api-order/internal/api/orders/models"
	"api-order/internal/entities"
	"context"
)

func (s *OrderServices) UpdateOrder(ctx context.Context, orderId string, req models.UpdateOrderRequest) (*entities.Orders, error) {
	err := s.repo.UpdateOrderStatusById(ctx, orderId, req.Status)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
