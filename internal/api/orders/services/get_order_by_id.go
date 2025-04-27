package services

import (
	"api-order/internal/entities"
	"context"
)

func (s *OrderServices) GetOrderById(ctx context.Context, orderId string) (*entities.Orders, error) {
	resp, err := s.repo.GetOrder(ctx, orderId)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
