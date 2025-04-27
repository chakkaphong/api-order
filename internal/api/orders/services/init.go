package services

import "api-order/internal/api/orders"

type OrderServices struct {
	repo orders.Repository
}

func NewServices(repo orders.Repository) orders.Service {
	return &OrderServices{
		repo: repo,
	}
}
