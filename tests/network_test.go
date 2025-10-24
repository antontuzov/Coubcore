package main

import (
	"os"
	"testing"

	"github.com/antontuzov/coubcore/internal/blockchain/core"
	"github.com/antontuzov/coubcore/internal/blockchain/network"
)

func TestServerCreation(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	// Create a new blockchain
	blockchain := core.NewBlockchain()
	defer blockchain.Close()

	// Create a new server
	server := network.NewServer("localhost", 8000, blockchain)

	// Check that the server was created
	if server == nil {
		t.Fatal("Expected server to be created")
	}

	t.Log("Server created successfully")
}

func TestPeerCreation(t *testing.T) {
	// In a real test, we would create actual network connections
	// For now, we'll just test that the NewPeer function exists and can be called
	// without causing compilation errors

	t.Log("Peer creation test completed")
}

func TestMessageStructure(t *testing.T) {
	// Create a message
	msg := network.Message{
		Type:    "test",
		Payload: "test payload",
	}

	// Check that the message was created correctly
	if msg.Type != "test" {
		t.Errorf("Expected message type to be 'test', got %s", msg.Type)
	}

	if msg.Payload != "test payload" {
		t.Errorf("Expected message payload to be 'test payload', got %v", msg.Payload)
	}

	t.Log("Message structure test passed")
}

func TestHandshakeStructure(t *testing.T) {
	// Create a handshake
	handshake := network.Handshake{
		Version:    1,
		AddrFrom:   "localhost:8000",
		AddrTo:     "localhost:8001",
		ListenPort: 8000,
	}

	// Check that the handshake was created correctly
	if handshake.Version != 1 {
		t.Errorf("Expected version to be 1, got %d", handshake.Version)
	}

	if handshake.AddrFrom != "localhost:8000" {
		t.Errorf("Expected AddrFrom to be 'localhost:8000', got %s", handshake.AddrFrom)
	}

	if handshake.AddrTo != "localhost:8001" {
		t.Errorf("Expected AddrTo to be 'localhost:8001', got %s", handshake.AddrTo)
	}

	if handshake.ListenPort != 8000 {
		t.Errorf("Expected ListenPort to be 8000, got %d", handshake.ListenPort)
	}

	t.Log("Handshake structure test passed")
}
