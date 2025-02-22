// service/metrics_service.go
package service

import (
	"github.com/shani34/order-management-system/metrics"
)

type MetricsService struct {
	Metrics *metrics.Metrics
}

func NewMetricsService(m *metrics.Metrics) *MetricsService {
	return &MetricsService{Metrics: m}
}

func (s *MetricsService) GetOrderMetrics() (map[string]interface{}, error) {
	return s.Metrics.GetMetrics()
}
