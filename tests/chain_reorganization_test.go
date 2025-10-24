package main

import (
	"os"
	"testing"

	"github.com/antontuzov/coubcore/internal/blockchain/core"
)

func TestChainReorganization(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	t.Run("ForkResolution", func(t *testing.T) {
		// Create two blockchains
		bc1 := core.NewBlockchain()
		bc2 := core.NewBlockchain()
		defer bc1.Close()
		defer bc2.Close()

		// Add blocks to both chains
		bc1.AddBlock("Block 1")
		bc1.AddBlock("Block 2")

		bc2.AddBlock("Block 1")
		bc2.AddBlock("Block 2")
		bc2.AddBlock("Block 3") // bc2 is longer

		// Get blocks from bc2
		blocks := bc2.GetBlocks()

		// Replace bc1's chain with bc2's chain
		success := bc1.ReplaceChain(blocks)

		// Check that the replacement was successful
		if !success {
			t.Error("Expected chain replacement to be successful")
		}

		// Check that bc1 now has the same length as bc2
		if bc1.Length() != bc2.Length() {
			t.Errorf("Expected bc1 length to be %d, got %d", bc2.Length(), bc1.Length())
		}

		// Check that the latest blocks are the same
		latest1 := bc1.GetLatestBlock()
		latest2 := bc2.GetLatestBlock()
		if latest1.Hash != latest2.Hash {
			t.Error("Expected latest blocks to be the same")
		}
	})

	t.Run("InvalidChainRejection", func(t *testing.T) {
		// Create two blockchains
		bc1 := core.NewBlockchain()
		bc2 := core.NewBlockchain()
		defer bc1.Close()
		defer bc2.Close()

		// Add blocks to bc2
		bc2.AddBlock("Block 1")
		bc2.AddBlock("Block 2")

		// Tamper with a block in bc2 to make it invalid
		block := bc2.GetBlockByIndex(1)
		block.Data = "Tampered Data"

		// Get blocks from bc2
		blocks := bc2.GetBlocks()

		// Try to replace bc1's chain with bc2's chain (should fail)
		success := bc1.ReplaceChain(blocks)

		// Check that the replacement failed
		if success {
			t.Error("Expected chain replacement to fail with invalid chain")
		}

		// Check that bc1 still has its original chain
		if bc1.Length() != 1 { // Only genesis block
			t.Errorf("Expected bc1 length to be 1, got %d", bc1.Length())
		}
	})

	t.Run("ShorterChainRejection", func(t *testing.T) {
		// Create two blockchains
		bc1 := core.NewBlockchain()
		bc2 := core.NewBlockchain()
		defer bc1.Close()
		defer bc2.Close()

		// Add more blocks to bc1
		bc1.AddBlock("Block 1")
		bc1.AddBlock("Block 2")
		bc1.AddBlock("Block 3")

		// Get blocks from bc2 (shorter chain)
		blocks := bc2.GetBlocks()

		// Try to replace bc1's chain with bc2's chain (should fail)
		success := bc1.ReplaceChain(blocks)

		// Check that the replacement failed
		if success {
			t.Error("Expected chain replacement to fail with shorter chain")
		}

		// Check that bc1 still has its original chain
		if bc1.Length() != 4 { // Genesis + 3 blocks
			t.Errorf("Expected bc1 length to be 4, got %d", bc1.Length())
		}
	})
}
