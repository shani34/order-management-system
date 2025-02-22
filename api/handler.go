// api/handlers.go
package api

import (
	"encoding/json"
	"net/http"
	"github.com/shani34/order-management-system/models"
	"github.com/shani34/order-management-system/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	Service *service.OrderService
	Metrics *service.MetricsService
}

func NewHandler(s *service.OrderService, m *service.MetricsService) *Handler {
	return &Handler{Service: s, Metrics: m}
}

// Create Order API 
func (h *Handler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Basic validation
	if order.UserID == 0 || len(order.ItemIDs) == 0 || order.TotalAmount <= 0 {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
		return
	}

	if err := h.Service.PlaceOrder(order); err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order received", "order_id": order.OrderID})
}

// Get Order Status API 
func (h *Handler) GetOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["order_id"]
	status, err := h.Service.GetOrderStatus(orderID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"order_id": orderID, "status": status})
}

// Get Metrics API
func (h *Handler) GetMetricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics, err := h.Metrics.GetOrderMetrics()
	if err != nil {
		http.Error(w, "Failed to fetch metrics", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(metrics)
}
