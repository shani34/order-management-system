package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/shani34/order-management-system/api"
	"github.com/shani34/order-management-system/config"
	"github.com/shani34/order-management-system/metrics"
	"github.com/shani34/order-management-system/models"
	"github.com/shani34/order-management-system/queue"
	"github.com/shani34/order-management-system/repository"
	"github.com/shani34/order-management-system/service"
)

func main() {
	db := config.ConnectDB()
	repo := repository.NewOrderRepository(db)
	orderQueue := queue.NewOrderQueue(1000)
	orderService := service.NewOrderService(repo, orderQueue)
	metricsService := service.NewMetricsService(metrics.NewMetrics(db))
	handler := api.NewHandler(orderService, metricsService)

	// Start order processing with multiple workers
	workerCount := 40
	go orderQueue.ProcessOrders(repo.UpdateOrderStatus, workerCount)

	// Define API routes
	r := mux.NewRouter()
	r.HandleFunc("/orders", handler.CreateOrderHandler).Methods("POST")
	r.HandleFunc("/orders/{order_id}/status", handler.GetOrderStatusHandler).Methods("GET")
	r.HandleFunc("/metrics", handler.GetMetricsHandler).Methods("GET") 

	simulateLoad(repo, orderQueue, 1000)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// simulateLoad generates 1,000 concurrent orders for scalability testing.
func simulateLoad(repo *repository.OrderRepository, orderQueue *queue.OrderQueue, orderCount int) {
	log.Println("Simulating", orderCount, "concurrent orders...")

	var wg sync.WaitGroup
	wg.Add(orderCount)

	startTime := time.Now()

	for i := 0; i < orderCount; i++ {
		go func(orderID string, userID int, totalAmount float64) {
			defer wg.Done()

			// Insert order into database
			err := repo.CreateOrder(models.Order{OrderID:orderID, UserID: userID,ItemIDs:[]int{101, 102, 103}, TotalAmount: totalAmount})
			if err != nil {
				log.Println("Error inserting order:", err)
				return
			}

			// Add order to processing queue
			orderQueue.AddToQueue(orderID)
		}(generateOrderID(i+1), rand.Intn(1000)+1, rand.Float64()*500)
	}

	wg.Wait()
	orderQueue.Wait() // Ensure all orders are processed

	totalTime := time.Since(startTime)
	log.Printf("All %d orders processed in %v\n", orderCount, totalTime)
}

// generateOrderID creates a unique order ID
func generateOrderID(index int) string {
	return fmt.Sprintf("ORD%d", index)
}
