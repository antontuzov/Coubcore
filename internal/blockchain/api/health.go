package api

import (
	"net/http"
	"time"

	"github.com/antontuzov/coubcore/internal/blockchain/core"
)

// HealthCheckResponse represents the response for health checks
type HealthCheckResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
	Chain     string    `json:"chain"`
	Height    int       `json:"height"`
	Peers     int       `json:"peers"`
}

// HealthCheckHandler handles health check requests
func HealthCheckHandler(blockchain *core.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get blockchain info
		height := blockchain.Length()

		// Create response
		response := HealthCheckResponse{
			Status:    "healthy",
			Timestamp: time.Now(),
			Version:   "1.0.0",
			Chain:     "main",
			Height:    height,
			Peers:     0, // In a real implementation, this would come from the network layer
		}

		// Send JSON response
		sendJSONResponse(w, response)
	}
}

// ReadyCheckHandler handles readiness check requests
func ReadyCheckHandler(blockchain *core.Blockchain) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if blockchain is valid
		isValid := blockchain.IsChainValid()

		if !isValid {
			http.Error(w, "Blockchain is not valid", http.StatusServiceUnavailable)
			return
		}

		// Create response
		response := HealthCheckResponse{
			Status:    "ready",
			Timestamp: time.Now(),
			Version:   "1.0.0",
			Chain:     "main",
			Height:    blockchain.Length(),
			Peers:     0, // In a real implementation, this would come from the network layer
		}

		// Send JSON response
		sendJSONResponse(w, response)
	}
}
