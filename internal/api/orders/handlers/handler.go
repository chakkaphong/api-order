package handlers

import (
	"api-order/configs"
	"api-order/internal/api/orders"
	"api-order/internal/api/orders/worker"
)

type OrdersHandler struct {
	service       orders.Service
	WorkerPool    *worker.WorkerPool
	serverConfigs configs.Server
}

func NewOrderHandler(service orders.Service, workerPool *worker.WorkerPool, serverConfigs configs.Server) *OrdersHandler {
	return &OrdersHandler{
		service:       service,
		WorkerPool:    workerPool,
		serverConfigs: serverConfigs,
	}
}
