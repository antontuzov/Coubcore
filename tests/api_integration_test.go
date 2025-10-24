package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/antontuzov/coubcore/internal/blockchain/api"
	"github.com/antontuzov/coubcore/internal/blockchain/core"
	"github.com/antontuzov/coubcore/internal/blockchain/network"
)

func TestAPIIntegration(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	// Create a new blockchain
	blockchain := core.NewBlockchain()
	defer blockchain.Close()

	// Create a new network server
	networkServer := network.NewServer("localhost", 8000, blockchain)

	// Create a new API server
	apiServer := api.NewServer(blockchain, networkServer, 8080)

	t.Run("GetBlockchainInfo", func(t *testing.T) {
		// Create a test request
		req, err := http.NewRequest("GET", "/api/v1/info", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Call the handler directly
		http.HandlerFunc(apiServer.GetInfo).ServeHTTP(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the content type
		expectedContentType := "application/json"
		if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
			t.Errorf("handler returned unexpected content type: got %v want %v",
				contentType, expectedContentType)
		}
	})

	t.Run("GetBlock", func(t *testing.T) {
		// Create a test request with query parameter
		req, err := http.NewRequest("GET", "/api/v1/block?index=0", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Call the handler directly
		http.HandlerFunc(apiServer.GetBlock).ServeHTTP(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	t.Run("GetTransaction", func(t *testing.T) {
		// Create a test request with query parameter
		req, err := http.NewRequest("GET", "/api/v1/transaction?txid=test", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Call the handler directly
		http.HandlerFunc(apiServer.GetTransaction).ServeHTTP(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	t.Run("GetBalance", func(t *testing.T) {
		// Create a test request with query parameter
		req, err := http.NewRequest("GET", "/api/v1/balance?address=test", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Call the handler directly
		http.HandlerFunc(apiServer.GetBalance).ServeHTTP(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	t.Run("GetPeers", func(t *testing.T) {
		// Create a test request
		req, err := http.NewRequest("GET", "/api/v1/peers", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Call the handler directly
		http.HandlerFunc(apiServer.GetPeers).ServeHTTP(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	t.Run("SendTransaction", func(t *testing.T) {
		// Create a test request
		req, err := http.NewRequest("GET", "/api/v1/send", nil)
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder
		rr := httptest.NewRecorder()

		// Call the handler directly
		http.HandlerFunc(apiServer.SendTransaction).ServeHTTP(rr, req)

		// Check the status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	t.Run("MissingParameters", func(t *testing.T) {
		// Test getBlock with missing index parameter
		req, err := http.NewRequest("GET", "/api/v1/block", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		http.HandlerFunc(apiServer.GetBlock).ServeHTTP(rr, req)

		// Should return bad request
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		// Test getTransaction with missing txid parameter
		req, err = http.NewRequest("GET", "/api/v1/transaction", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr = httptest.NewRecorder()
		http.HandlerFunc(apiServer.GetTransaction).ServeHTTP(rr, req)

		// Should return bad request
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		// Test getBalance with missing address parameter
		req, err = http.NewRequest("GET", "/api/v1/balance", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr = httptest.NewRecorder()
		http.HandlerFunc(apiServer.GetBalance).ServeHTTP(rr, req)

		// Should return bad request
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})
}
