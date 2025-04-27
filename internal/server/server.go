package server

import (
	"api-order/configs"
	worker "api-order/internal/api/orders/worker"
	"api-order/internal/database"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	port       int
	postgres   database.Service
	echo       *echo.Echo
	workerPool *worker.WorkerPool
}

func NewServer() (*http.Server, *worker.WorkerPool) {
	cfg := configs.GetConfig()

	postgres := database.New()

	NewServer := &Server{
		port:     cfg.App.Port,
		postgres: postgres,
	}

	// Setup Echo routes and worker pool
	e, workerPool := NewServer.RegisterRoutes()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      e,
		IdleTimeout:  cfg.Server.IdleTimeout,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.ReadTimeout,
	}

	log.Default().Printf("Server is running, port:%d", NewServer.port)

	return server, workerPool
}
