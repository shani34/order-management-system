package queue

import (
	"log"
	"sync"
	"time"
)

type OrderQueue struct {
	Orders chan string
	Wg     sync.WaitGroup
}

// NewOrderQueue initializes an order queue with a buffer size.
func NewOrderQueue(size int) *OrderQueue {
	return &OrderQueue{Orders: make(chan string, size)}
}

// AddToQueue adds an order to the queue.
func (q *OrderQueue) AddToQueue(orderID string) {
	q.Wg.Add(1)
	q.Orders <- orderID
}

// ProcessOrders starts multiple workers to process orders concurrently.
func (q *OrderQueue) ProcessOrders(updateStatusFunc func(orderID string, status string) error, workers int) {
	for i := 0; i < workers; i++ {
		go func(workerID int) {
			for orderID := range q.Orders {
				log.Printf("[Worker %d] Processing order: %s\n", workerID, orderID)
				
				_ = updateStatusFunc(orderID, "Processing")
				time.Sleep(3 * time.Second) // Simulate processing delay
				
				_ = updateStatusFunc(orderID, "Completed")
				log.Printf("[Worker %d] Order completed: %s\n", workerID, orderID)
				
				q.Wg.Done()
			}
		}(i + 1)
	}
}

// Wait waits for all queued orders to be processed.
func (q *OrderQueue) Wait() {
	q.Wg.Wait()
}
