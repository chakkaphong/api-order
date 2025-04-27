package handlers

import (
	"api-order/internal/api/orders/models"
	"api-order/internal/global/responses"
	"api-order/internal/global/utils"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *OrdersHandler) GetOrders(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), h.serverConfigs.ReadTimeout)
	defer cancel()

	paginate := utils.GetPaginationQuery(c)

	resp, totalPage, err := h.service.GetOrders(ctx, paginate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	orderResp := []models.OrderResponse{}

	for _, order := range resp {
		orderResp = append(orderResp, *h.buildOrderData(&order))
	}

	paginate.Total = totalPage

	return c.JSON(http.StatusOK, responses.SuccessPaging(orderResp, paginate, c))
}
