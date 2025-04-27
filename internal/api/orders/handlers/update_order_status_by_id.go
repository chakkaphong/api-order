package handlers

import (
	"api-order/internal/api/orders/models"
	"api-order/internal/global/responses"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *OrdersHandler) UpdateOrder(c echo.Context) error {
	orderId := c.Param("id")
	var req models.UpdateOrderRequest
	ctx, cancel := context.WithTimeout(c.Request().Context(), h.serverConfigs.WriteTimeout)
	defer cancel()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequest(err.Error()))
	}

	_, err := h.service.UpdateOrder(ctx, orderId, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, responses.Success("order updated"))
}
