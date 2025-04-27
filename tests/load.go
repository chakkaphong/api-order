// tests/load_test.go

package main

import (
	"api-order/internal/api/orders/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

const (
	url            = "http://localhost:9001/v1/orders"
	numRequests    = 1000
	concurrentReqs = 10
)

func main() {
	var wg sync.WaitGroup
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	// Start the load test
	startTime := time.Now()

	for i := 0; i < concurrentReqs; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < numRequests; j++ {
				orderReq := models.CreateOrderRequest{
					CustomerName: fmt.Sprintf("Customer-%d", j),
					OrderItems: []models.OrderItem{
						{ProductName: "Product A", Quantity: 2, Price: decimal.NewFromFloat(23.1)},
						{ProductName: "Product B", Quantity: 1, Price: decimal.NewFromFloat(7.1)},
					},
				}

				orderJSON, err := json.Marshal(orderReq)
				if err != nil {
					log.Printf("Error marshalling request: %v\n", err)
					return
				}

				resp, err := client.Post(url, "application/json", bytes.NewBuffer(orderJSON))
				if err != nil {
					log.Printf("Error sending request: %v\n", err)
					return
				}

				resp.Body.Close()

				if resp.StatusCode != http.StatusAccepted {
					log.Printf("Worker %d: Order request failed with status code %d\n", workerID, resp.StatusCode)
				}
			}
		}(i)
	}

	wg.Wait()

	duration := time.Since(startTime)
	fmt.Printf("Load Test Completed in %v\n", duration)
}
