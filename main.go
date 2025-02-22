package main

import (
	"log"
	"net/http"
	"github.com/shani34/order-management-system/api"
	"github.com/shani34/order-management-system/config"
	"github.com/shani34/order-management-system/queue"
	"github.com/shani34/order-management-system/repository"
	"github.com/shani34/order-management-system/service"
	"github.com/shani34/order-management-system/metrics"


	"github.com/gorilla/mux"
)

func main() {
	db := config.ConnectDB()
	repo := repository.NewOrderRepository(db)
	orderQueue := queue.NewOrderQueue(1000)
	orderService := service.NewOrderService(repo, orderQueue)
	metricsService := service.NewMetricsService(metrics.NewMetrics(db))
	handler := api.NewHandler(orderService, metricsService)

	// Start order processing in a separate goroutine
	go orderQueue.ProcessOrders(repo.UpdateOrderStatus)

	// Define API routes
	r := mux.NewRouter()
	r.HandleFunc("/orders", handler.CreateOrderHandler).Methods("POST")
	r.HandleFunc("/orders/{order_id}/status", handler.GetOrderStatusHandler).Methods("GET")
	r.HandleFunc("/metrics", handler.GetMetricsHandler).Methods("GET") // NEW

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}