package main

import (
	"os"
	"testing"

	"github.com/antontuzov/coubcore/internal/blockchain/core"
)

func TestBlockValidationComprehensive(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	// Test valid block
	t.Run("ValidBlock", func(t *testing.T) {
		// Create a previous block
		prevBlock := core.NewBlock(0, "", "genesis data")

		// Create a new block
		block := core.NewBlock(1, prevBlock.Hash, "test data")

		// Validate the block against the previous block
		if !block.Validate(prevBlock) {
			t.Error("Expected block to be valid")
		}
	})

	// Test invalid block (wrong previous hash)
	t.Run("InvalidPreviousHash", func(t *testing.T) {
		// Create a previous block
		prevBlock := core.NewBlock(0, "", "genesis data")

		// Create a new block
		block := core.NewBlock(1, "wrongHash", "test data")

		// Validate the block against the previous block
		if block.Validate(prevBlock) {
			t.Error("Expected block to be invalid")
		}
	})

	// Test invalid block (wrong index)
	t.Run("InvalidIndex", func(t *testing.T) {
		// Create a previous block
		prevBlock := core.NewBlock(0, "", "genesis data")

		// Create a new block with wrong index
		block := core.NewBlock(2, prevBlock.Hash, "test data")

		// Validate the block against the previous block
		if block.Validate(prevBlock) {
			t.Error("Expected block to be invalid")
		}
	})

	// Test invalid block (tampered data)
	t.Run("TamperedData", func(t *testing.T) {
		// Create a previous block
		prevBlock := core.NewBlock(0, "", "genesis data")

		// Create a new block
		block := core.NewBlock(1, prevBlock.Hash, "test data")

		// Tamper with the block data
		block.Data = "tampered data"

		// Validate the block (should fail because hash doesn't match)
		if block.Validate(prevBlock) {
			t.Error("Expected block to be invalid")
		}
	})

	// Test invalid block (wrong timestamp)
	t.Run("InvalidTimestamp", func(t *testing.T) {
		// Create a previous block
		prevBlock := core.NewBlock(0, "", "genesis data")

		// Create a new block
		block := core.NewBlock(1, prevBlock.Hash, "test data")

		// Set the timestamp to be before the previous block
		block.Timestamp = prevBlock.Timestamp.Add(-1 * 1000000000) // 1 second before

		// Validate the block against the previous block
		if block.Validate(prevBlock) {
			t.Error("Expected block to be invalid")
		}
	})
}
