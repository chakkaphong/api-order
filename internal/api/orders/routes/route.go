package routes

import (
	"api-order/configs"
	"api-order/internal/api/orders"
	"api-order/internal/api/orders/handlers"
	"api-order/internal/api/orders/models"
	"api-order/internal/api/orders/worker"

	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
)

func RegisterProductRoutes(rootAPI echoswagger.ApiRoot, service orders.Service, workerPool *worker.WorkerPool) {
	cfg := configs.GetConfig()
	controller := handlers.NewOrderHandler(service, workerPool, cfg.Server)

	group := rootAPI.Group("order", "v1/orders").SetDescription("Order API")

	group.POST("", controller.PostOrder).
		SetSummary("create new orders").
		SetDescription("create new orders").
		AddParamBody(models.CreateOrderRequest{}, "body", "send request to create order", true).
		SetResponseContentType(echo.MIMEApplicationJSON)

	group.GET("", controller.GetOrders).
		SetSummary("get order list").
		SetDescription("get order list")

	group.GET("/:id", controller.GetOrderById).
		SetSummary("get order by id").
		SetDescription("get order by id")

	group.PUT("/:id/status", controller.UpdateOrder).
		SetSummary("update order by id").
		SetDescription("update order by id")
}
