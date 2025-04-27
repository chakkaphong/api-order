package handlers

import (
	"api-order/internal/api/orders/models"
	"api-order/internal/entities"
	"api-order/internal/global/responses"
	"context"
	"errors"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (h *OrdersHandler) GetOrderById(c echo.Context) error {
	orderId := c.Param("id")

	ctx, cancel := context.WithTimeout(c.Request().Context(), h.serverConfigs.ReadTimeout)
	defer cancel()

	resp, err := h.service.GetOrderById(ctx, orderId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, responses.DataNotFound("Order was not found"))
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	orderResp := h.buildOrderData(resp)

	return c.JSON(http.StatusOK, responses.Success(orderResp))
}

func (h *OrdersHandler) buildOrderData(order *entities.Orders) *models.OrderResponse {
	resp := models.OrderResponse{
		Id:           order.Id,
		CustomerName: order.CustomerName,
		TotalAmount:  order.TotalAmount,
		Status:       order.Status,
		CreatedAt:    order.CreatedAt,
	}

	for _, item := range order.OrderItems {
		dto := models.OrderItemsResponse{}
		copier.Copy(&dto, &item)

		resp.OrderItems = append(resp.OrderItems, dto)
	}

	return &resp
}
