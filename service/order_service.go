package service



import (
	"github.com/shani34/order-management-system/models"
	"github.com/shani34/order-management-system/queue"
	"github.com/shani34/order-management-system/repository"
)

type OrderService struct {
	Repo  *repository.OrderRepository
	Queue *queue.OrderQueue
}

func NewOrderService(repo *repository.OrderRepository, q *queue.OrderQueue) *OrderService {
	return &OrderService{Repo: repo, Queue: q}
}

func (s *OrderService) PlaceOrder(order models.Order) error {
	order.Status = "Pending"
	err := s.Repo.CreateOrder(order)
	if err != nil {
		return err
	}
	s.Queue.AddToQueue(order.OrderID)
	return nil
}

func (s *OrderService) GetOrderStatus(orderID string) (string, error) {
	return s.Repo.GetOrderStatus(orderID)
}
