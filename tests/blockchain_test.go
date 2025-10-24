package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/antontuzov/coubcore/internal/blockchain/core"
)

func TestBlockchainFunctionality(t *testing.T) {
	fmt.Println("Testing blockchain functionality...")

	// Remove existing database for clean test
	os.Remove("blockchain.db")

	// Create a new blockchain
	bc := core.NewBlockchain()
	defer bc.Close()

	// Check that the blockchain was created
	if bc == nil {
		t.Fatal("Expected blockchain to be created")
	}

	// Check that the genesis block was created
	if bc.Length() != 1 {
		t.Errorf("Expected blockchain length to be 1, got %d", bc.Length())
	}

	// Add a new block
	newBlock := bc.AddBlock("Test Data")

	// Check that the block was added
	if bc.Length() != 2 {
		t.Errorf("Expected blockchain length to be 2, got %d", bc.Length())
	}

	// Check that the new block is correct
	if newBlock.Index != 1 {
		t.Errorf("Expected new block index to be 1, got %d", newBlock.Index)
	}

	if newBlock.Data != "Test Data" {
		t.Errorf("Expected new block data to be 'Test Data', got %s", newBlock.Data)
	}

	// Check that the chain is valid
	if !bc.IsChainValid() {
		t.Error("Expected blockchain to be valid")
	}

	fmt.Println("All tests passed!")
}

func TestBlockchainCreation(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	// Create a new blockchain
	bc := core.NewBlockchain()
	defer bc.Close()

	// Check that the blockchain was created
	if bc == nil {
		t.Fatal("Expected blockchain to be created")
	}

	// Check that the genesis block was created
	if bc.Length() != 1 {
		t.Errorf("Expected blockchain length to be 1, got %d", bc.Length())
	}

	// Check that the genesis block is valid
	genesisBlock := bc.GetBlockByIndex(0)
	if genesisBlock == nil {
		t.Fatal("Expected genesis block to exist")
	}

	if genesisBlock.Index != 0 {
		t.Errorf("Expected genesis block index to be 0, got %d", genesisBlock.Index)
	}

	if genesisBlock.Data != "Genesis Block" {
		t.Errorf("Expected genesis block data to be 'Genesis Block', got %s", genesisBlock.Data)
	}
}

func TestAddBlock(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	// Create a new blockchain
	bc := core.NewBlockchain()
	defer bc.Close()

	// Add a new block
	newBlock := bc.AddBlock("Test Data")

	// Check that the block was added
	if bc.Length() != 2 {
		t.Errorf("Expected blockchain length to be 2, got %d", bc.Length())
	}

	// Check that the new block is correct
	if newBlock.Index != 1 {
		t.Errorf("Expected new block index to be 1, got %d", newBlock.Index)
	}

	if newBlock.Data != "Test Data" {
		t.Errorf("Expected new block data to be 'Test Data', got %s", newBlock.Data)
	}

	if newBlock.PreviousHash != bc.GetBlockByIndex(0).Hash {
		t.Errorf("Expected new block previousHash to match genesis block hash")
	}
}

func TestGetLatestBlock(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	// Create a new blockchain
	bc := core.NewBlockchain()
	defer bc.Close()

	// Add a new block
	bc.AddBlock("Test Data")

	// Get the latest block
	latestBlock := bc.GetLatestBlock()

	// Check that the latest block is correct
	if latestBlock.Index != 1 {
		t.Errorf("Expected latest block index to be 1, got %d", latestBlock.Index)
	}

	if latestBlock.Data != "Test Data" {
		t.Errorf("Expected latest block data to be 'Test Data', got %s", latestBlock.Data)
	}
}

func TestIsChainValid(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	// Create a new blockchain
	bc := core.NewBlockchain()
	defer bc.Close()

	// Add some blocks
	bc.AddBlock("Block 1")
	bc.AddBlock("Block 2")
	bc.AddBlock("Block 3")

	// Check that the chain is valid
	if !bc.IsChainValid() {
		t.Error("Expected blockchain to be valid")
	}

	// Tamper with a block to make it invalid
	block := bc.GetBlockByIndex(1)
	block.Data = "Tampered Data"

	// Check that the chain is now invalid
	if bc.IsChainValid() {
		t.Error("Expected blockchain to be invalid after tampering")
	}
}

func TestReplaceChain(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	// Create two blockchains
	bc1 := core.NewBlockchain()
	bc2 := core.NewBlockchain()
	defer bc1.Close()
	defer bc2.Close()

	// Add more blocks to bc2
	bc2.AddBlock("Block 1")
	bc2.AddBlock("Block 2")

	// Replace bc1's chain with bc2's chain
	blocks := bc2.GetBlocks()
	success := bc1.ReplaceChain(blocks)

	// Check that the replacement was successful
	if !success {
		t.Error("Expected chain replacement to be successful")
	}

	// Check that bc1 now has the same length as bc2
	if bc1.Length() != bc2.Length() {
		t.Errorf("Expected bc1 length to be %d, got %d", bc2.Length(), bc1.Length())
	}

	// Try to replace with a shorter chain (should fail)
	success = bc2.ReplaceChain(bc1.GetBlocks()[:2])
	if success {
		t.Error("Expected chain replacement to fail with shorter chain")
	}
}
