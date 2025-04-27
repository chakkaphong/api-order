package server

import (
	"api-order/configs"

	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"

	ordersRepository "api-order/internal/api/orders/repo"
	ordersRoutes "api-order/internal/api/orders/routes"
	ordersService "api-order/internal/api/orders/services"
	worker "api-order/internal/api/orders/worker"
)

func (s *Server) RegisterRoutes() (*echo.Echo, *worker.WorkerPool) {
	cfg := configs.GetConfig()
	e := echo.New()
	var apiRoot echoswagger.ApiRoot

	apiRoot = echoswagger.New(e, "docs", &echoswagger.Info{
		Title:       cfg.App.Name,
		Description: "Template API Documentation",
		Version:     "0.0.1",
	})

	// Init modules
	orderRepo := ordersRepository.NewOrderRepository(s.postgres)
	orderSrv := ordersService.NewServices(orderRepo)
	workerPool := worker.NewWorkerPool(orderSrv, cfg.Server.NumWorker, cfg.Server.QueueSize)

	// Register Routes
	ordersRoutes.RegisterProductRoutes(apiRoot, orderSrv, workerPool)

	workerPool.Start()

	return e, workerPool
}
