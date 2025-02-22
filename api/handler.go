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
}

func NewHandler(s *service.OrderService) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.Service.PlaceOrder(order); err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Order received", "order_id": order.OrderID})
}

func (h *Handler) GetOrderStatusHandler(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["order_id"]
	status, err := h.Service.GetOrderStatus(orderID)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"order_id": orderID, "status": status})
}
