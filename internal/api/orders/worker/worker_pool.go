package worker

import (
	"api-order/internal/api/orders"
	"api-order/internal/api/orders/models"
	"context"
	"sync"
	"time"
)

type OrderJob struct {
	Ctx   context.Context
	Order models.CreateOrderRequest
}

type WorkerPool struct {
	Queue        chan OrderJob
	Wg           *sync.WaitGroup
	OrderService orders.Service
	NumWorker    int
}

func NewWorkerPool(service orders.Service, numWorker int, queueSize int) *WorkerPool {
	return &WorkerPool{
		Queue:        make(chan OrderJob, queueSize),
		Wg:           &sync.WaitGroup{},
		OrderService: service,
		NumWorker:    numWorker,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.NumWorker; i++ {
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(workerID int) {
	for job := range wp.Queue {
		jobCtx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()

		wp.OrderService.CreateOrder(jobCtx, job.Order)
		wp.Wg.Done()
	}
}

func (wp *WorkerPool) Shutdown() {
	close(wp.Queue)
	wp.Wg.Wait()
}
