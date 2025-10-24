package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MetricsHandler handles Prometheus metrics requests
func MetricsHandler() http.HandlerFunc {
	return promhttp.Handler().ServeHTTP
}
