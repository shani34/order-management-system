
package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/shani34/order-management-system/models"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) CreateOrder(order models.Order) error {
	itemIDs, _ := json.Marshal(order.ItemIDs)
	_, err := r.DB.Exec("INSERT INTO orders (order_id, user_id, item_ids, total_amount, status) VALUES (?, ?, ?, ?, ?)",
		order.OrderID, order.UserID, string(itemIDs), order.TotalAmount, order.Status)
	return err
}

func (r *OrderRepository) GetOrderStatus(orderID string) (string, error) {
	var status string
	err := r.DB.QueryRow("SELECT status FROM orders WHERE order_id = ?", orderID).Scan(&status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("order not found")
		}
		return "", err
	}
	return status, nil
}

func (r *OrderRepository) UpdateOrderStatus(orderID string, status string) error {
	_, err := r.DB.Exec("UPDATE orders SET status = ? WHERE order_id = ?", status, orderID)
	return err
}
