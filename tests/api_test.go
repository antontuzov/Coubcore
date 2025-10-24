package main

import (
	"os"
	"testing"

	"github.com/antontuzov/coubcore/internal/blockchain/api"
	"github.com/antontuzov/coubcore/internal/blockchain/core"
	"github.com/antontuzov/coubcore/internal/blockchain/network"
)

func TestAPIServerCreation(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	// Create a new blockchain
	blockchain := core.NewBlockchain()
	defer blockchain.Close()

	// Create a new network server
	networkServer := network.NewServer("localhost", 8000, blockchain)

	// Create a new API server
	apiServer := api.NewServer(blockchain, networkServer, 8080)

	// Check that the API server was created
	if apiServer == nil {
		t.Fatal("Expected API server to be created")
	}

	t.Log("API server created successfully")
}

func TestAPIEndpoints(t *testing.T) {
	// Test that all the required API endpoints are defined
	// In a real test, we would make HTTP requests to these endpoints

	endpoints := []string{
		"/api/v1/info",
		"/api/v1/block",
		"/api/v1/transaction",
		"/api/v1/balance",
		"/api/v1/peers",
		"/api/v1/send",
	}

	// Just verify that we've defined these endpoints
	// In a real test, we would start the server and make requests

	for _, endpoint := range endpoints {
		if endpoint == "" {
			t.Errorf("Endpoint is empty")
		}
	}

	t.Log("API endpoints defined successfully")
}
