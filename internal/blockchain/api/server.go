package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/antontuzov/coubcore/internal/blockchain/core"
	"github.com/antontuzov/coubcore/internal/blockchain/network"
)

// Server represents the JSON-RPC API server
type Server struct {
	blockchain *core.Blockchain
	network    *network.Server
	port       int
}

// NewServer creates a new API server
func NewServer(blockchain *core.Blockchain, network *network.Server, port int) *Server {
	return &Server{
		blockchain: blockchain,
		network:    network,
		port:       port,
	}
}

// Start starts the JSON-RPC API server
func (s *Server) Start() error {
	// Register API endpoints
	http.HandleFunc("/api/v1/info", s.GetInfo)
	http.HandleFunc("/api/v1/block", s.GetBlock)
	http.HandleFunc("/api/v1/transaction", s.GetTransaction)
	http.HandleFunc("/api/v1/balance", s.GetBalance)
	http.HandleFunc("/api/v1/peers", s.GetPeers)
	http.HandleFunc("/api/v1/send", s.SendTransaction)

	// Health check endpoints
	http.HandleFunc("/health", HealthCheckHandler(s.blockchain))
	http.HandleFunc("/ready", ReadyCheckHandler(s.blockchain))

	// Metrics endpoint
	http.Handle("/metrics", MetricsHandler())

	log.Printf("JSON-RPC API server started on port %d", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
}

// GetInfo returns blockchain information
func (s *Server) GetInfo(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	info := map[string]interface{}{
		"length": s.blockchain.Length(),
		"latest": s.blockchain.GetLatestBlock().Index,
	}

	json.NewEncoder(w).Encode(info)
}

// GetBlock returns block information
func (s *Server) GetBlock(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Get the block index from query parameters
	index := r.URL.Query().Get("index")
	if index == "" {
		http.Error(w, "Missing index parameter", http.StatusBadRequest)
		return
	}

	// TODO: Parse index and retrieve block
	// For now, we'll return a placeholder response
	response := map[string]interface{}{
		"index": index,
		"data":  "block data",
	}

	json.NewEncoder(w).Encode(response)
}

// GetTransaction returns transaction information
func (s *Server) GetTransaction(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Get the transaction ID from query parameters
	txid := r.URL.Query().Get("txid")
	if txid == "" {
		http.Error(w, "Missing txid parameter", http.StatusBadRequest)
		return
	}

	// TODO: Retrieve transaction by ID
	// For now, we'll return a placeholder response
	response := map[string]interface{}{
		"txid": txid,
		"data": "transaction data",
	}

	json.NewEncoder(w).Encode(response)
}

// GetBalance returns the balance of an address
func (s *Server) GetBalance(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Get the address from query parameters
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Missing address parameter", http.StatusBadRequest)
		return
	}

	// TODO: Calculate balance for the address
	// For now, we'll return a placeholder response
	response := map[string]interface{}{
		"address": address,
		"balance": 0,
	}

	json.NewEncoder(w).Encode(response)
}

// GetPeers returns the list of connected peers
func (s *Server) GetPeers(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// TODO: Get the list of connected peers
	// For now, we'll return a placeholder response
	response := map[string]interface{}{
		"peers": []string{},
	}

	json.NewEncoder(w).Encode(response)
}

// SendTransaction handles sending a transaction
func (s *Server) SendTransaction(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// TODO: Handle transaction submission
	// For now, we'll return a placeholder response
	response := map[string]interface{}{
		"status": "success",
		"txid":   "transaction_id",
	}

	json.NewEncoder(w).Encode(response)
}
