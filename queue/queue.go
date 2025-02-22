package queue

import (
	"log"
	"time"
)

type OrderQueue struct {
	Orders chan string
}

func NewOrderQueue(size int) *OrderQueue {
	return &OrderQueue{Orders: make(chan string, size)}
}

func (q *OrderQueue) AddToQueue(orderID string) {
	q.Orders <- orderID
}

func (q *OrderQueue) ProcessOrders(updateStatusFunc func(orderID string, status string) error) {
	for orderID := range q.Orders {
		log.Println("Processing order:", orderID)
		_ = updateStatusFunc(orderID, "Processing")
		time.Sleep(3 * time.Second) // Simulate processing time
		_ = updateStatusFunc(orderID, "Completed")
		log.Println("Order completed:", orderID)
	}
}
