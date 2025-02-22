package main

import (
	"log"
	"net/http"
	"github.com/shani34/order-management-system/api"
	"github.com/shani34/order-management-system/config"
	"github.com/shani34/order-management-system/queue"
	"github.com/shani34/order-management-system/repository"
	"github.com/shani34/order-management-system/service"

	"github.com/gorilla/mux"
)

func main() {
	db := config.ConnectDB()
	repo := repository.NewOrderRepository(db)
	orderQueue := queue.NewOrderQueue(1000)
	service := service.NewOrderService(repo, orderQueue)
	handler := api.NewHandler(service)

	// Start processing queue in a separate goroutine
	go orderQueue.ProcessOrders(repo.UpdateOrderStatus)

	r := mux.NewRouter()
	r.HandleFunc("/orders", handler.CreateOrderHandler).Methods("POST")
	r.HandleFunc("/orders/{order_id}/status", handler.GetOrderStatusHandler).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
