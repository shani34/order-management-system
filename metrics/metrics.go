package metrics

import (
	"database/sql"
	"log"
)

type Metrics struct {
	DB *sql.DB
}

func NewMetrics(db *sql.DB) *Metrics {
	return &Metrics{DB: db}
}

func (m *Metrics) GetMetrics() (map[string]interface{}, error) {
	var totalOrders int
	var avgProcessingTime float64
	var pending, processing, completed int

	// Get total orders
	err := m.DB.QueryRow("SELECT COUNT(*) FROM orders").Scan(&totalOrders)
	if err != nil {
		log.Println("Error fetching total orders:", err)
		return nil, err
	}

	// Get average processing time
	err = m.DB.QueryRow(`
		SELECT COALESCE(AVG(julianday(updated_at) - julianday(created_at)) * 86400, 0) FROM orders WHERE status = 'Completed'
	`).Scan(&avgProcessingTime)
	if err != nil {
		log.Println("Error fetching average processing time:", err)
		return nil, err
	}

	// Get count of each status
	err = m.DB.QueryRow("SELECT COUNT(*) FROM orders WHERE status = 'Pending'").Scan(&pending)
	err = m.DB.QueryRow("SELECT COUNT(*) FROM orders WHERE status = 'Processing'").Scan(&processing)
	err = m.DB.QueryRow("SELECT COUNT(*) FROM orders WHERE status = 'Completed'").Scan(&completed)
	if err != nil {
		log.Println("Error fetching order status counts:", err)
		return nil, err
	}

	metrics := map[string]interface{}{
		"total_orders":          totalOrders,
		"average_processing_time": avgProcessingTime,
		"order_status_counts": map[string]int{
			"pending":    pending,
			"processing": processing,
			"completed":  completed,
		},
	}

	return metrics, nil
}

