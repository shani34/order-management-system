package models


import "time"

type Order struct {
	OrderID     string    `json:"order_id"`
	UserID      int       `json:"user_id"`
	ItemIDs     []int     `json:"item_ids"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
