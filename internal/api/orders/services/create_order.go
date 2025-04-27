package services

import (
	"api-order/internal/api/orders"
	"api-order/internal/api/orders/models"
	"api-order/internal/entities"
	"context"

	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

func (s *OrderServices) CreateOrder(ctx context.Context, req models.CreateOrderRequest) error {
	if err := s.repo.WithTransaction(ctx, func(rp orders.Repository) error {

		totalAmount := lo.Reduce(req.OrderItems, func(agg decimal.Decimal, item models.OrderItem, _ int) decimal.Decimal {
			return agg.Add(item.Price)
		}, decimal.Zero)

		order := entities.Orders{
			CustomerName: req.CustomerName,
			Status:       entities.OrderStatusEnumCreated,
			TotalAmount:  totalAmount,
		}
		createOrderResp, err := rp.SaveOrder(ctx, order)
		if err != nil {
			return err
		}

		orderItems := []entities.OrderItems{}
		for _, item := range req.OrderItems {
			orderItems = append(orderItems, entities.OrderItems{
				OrderId:     createOrderResp.Id,
				ProductName: item.ProductName,
				Quantity:    item.Quantity,
				Price:       item.Price,
			})
		}

		_, err = rp.SaveOrderItem(ctx, orderItems)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
