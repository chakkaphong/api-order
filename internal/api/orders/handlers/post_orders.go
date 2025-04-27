package handlers

import (
	"context"
	"net/http"
	"time"

	"api-order/internal/api/orders/models"
	"api-order/internal/api/orders/worker"
	"api-order/internal/global/responses"

	"github.com/labstack/echo/v4"
)

func (h *OrdersHandler) PostOrder(c echo.Context) error {
	var req models.CreateOrderRequest

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Minute) // Adjust timeout as needed
	defer cancel()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequest(err.Error()))
	}

	h.WorkerPool.Wg.Add(1)

	// Before passing context to worker, check if it's already canceled
	if ctx.Err() != nil {
		return c.JSON(http.StatusServiceUnavailable, responses.BadRequest("context canceled before processing"))
	}

	select {
	case h.WorkerPool.Queue <- worker.OrderJob{Ctx: ctx, Order: req}:
		return c.JSON(http.StatusAccepted, responses.Success("order accepted"))
	default:
		h.WorkerPool.Wg.Done()
		return c.JSON(http.StatusCreated, responses.Success("server busy, try again later"))
	}
}
