package main

import (
	"testing"

	"github.com/antontuzov/coubcore/internal/blockchain/core"
)

func TestBlockCreation(t *testing.T) {
	// Create a new block
	block := core.NewBlock(1, "previousHash", "test data")

	// Check that the block was created correctly
	if block.Index != 1 {
		t.Errorf("Expected index 1, got %d", block.Index)
	}

	if block.PreviousHash != "previousHash" {
		t.Errorf("Expected previousHash, got %s", block.PreviousHash)
	}

	if block.Data != "test data" {
		t.Errorf("Expected test data, got %s", block.Data)
	}

	// Check that the hash is not empty
	if block.Hash == "" {
		t.Error("Expected hash to be calculated")
	}
}

func TestBlockHashCalculation(t *testing.T) {
	// Create a new block
	block := core.NewBlock(1, "previousHash", "test data")

	// Calculate the hash manually
	calculatedHash := block.CalculateHash()

	// Check that the calculated hash matches the block's hash
	if block.Hash != calculatedHash {
		t.Errorf("Expected hash %s, got %s", calculatedHash, block.Hash)
	}
}

func TestBlockSerialization(t *testing.T) {
	// Create a new block
	block := core.NewBlock(1, "previousHash", "test data")

	// Serialize the block
	data, err := block.Serialize()
	if err != nil {
		t.Fatalf("Error serializing block: %v", err)
	}

	// Check that we got some data
	if len(data) == 0 {
		t.Error("Expected serialized data")
	}

	// Deserialize the block
	deserializedBlock, err := core.Deserialize(data)
	if err != nil {
		t.Fatalf("Error deserializing block: %v", err)
	}

	// Check that the deserialized block matches the original
	if deserializedBlock.Index != block.Index {
		t.Errorf("Expected index %d, got %d", block.Index, deserializedBlock.Index)
	}

	if deserializedBlock.Hash != block.Hash {
		t.Errorf("Expected hash %s, got %s", block.Hash, deserializedBlock.Hash)
	}
}

func TestBlockValidation(t *testing.T) {
	// Create a previous block
	prevBlock := core.NewBlock(0, "", "genesis data")

	// Create a new block
	block := core.NewBlock(1, prevBlock.Hash, "test data")

	// Validate the block against the previous block
	if !block.Validate(prevBlock) {
		t.Error("Expected block to be valid")
	}

	// Test invalid block (wrong previous hash)
	invalidBlock := core.NewBlock(1, "wrongHash", "test data")
	if invalidBlock.Validate(prevBlock) {
		t.Error("Expected block to be invalid")
	}

	// Test invalid block (wrong index)
	invalidBlock2 := core.NewBlock(2, prevBlock.Hash, "test data")
	if invalidBlock2.Validate(prevBlock) {
		t.Error("Expected block to be invalid")
	}
}
