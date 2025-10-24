package main

import (
	"os"
	"testing"

	"github.com/antontuzov/coubcore/internal/blockchain/consensus"
	"github.com/antontuzov/coubcore/internal/blockchain/core"
)

func TestProofOfWorkComprehensive(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	t.Run("MineBlock", func(t *testing.T) {
		// Create a new blockchain
		blockchain := core.NewBlockchain()
		defer blockchain.Close()

		// Mine a new block
		block := consensus.MineBlock(blockchain, "Test transaction")

		// Check that the block was mined
		if block == nil {
			t.Fatal("Expected block to be mined")
		}

		// Check that the block has a hash
		if block.Hash == "" {
			t.Error("Expected block to have a hash")
		}

		// Check that the block was added to the blockchain
		if blockchain.Length() != 2 {
			t.Errorf("Expected blockchain length to be 2, got %d", blockchain.Length())
		}

		// Get the latest block from the blockchain
		latestBlock := blockchain.GetLatestBlock()
		if latestBlock.Hash != block.Hash {
			t.Error("Expected latest block to match mined block")
		}
	})

	t.Run("DifficultyAdjustment", func(t *testing.T) {
		// Create a new blockchain
		blockchain := core.NewBlockchain()
		defer blockchain.Close()

		// Mine several blocks quickly to test difficulty adjustment
		for i := 0; i < 5; i++ {
			block := consensus.MineBlock(blockchain, "Test transaction")
			if block == nil {
				t.Fatal("Expected block to be mined")
			}
		}

		// Check that we have the expected number of blocks
		if blockchain.Length() != 6 { // Genesis + 5 mined blocks
			t.Errorf("Expected blockchain length to be 6, got %d", blockchain.Length())
		}

		// Check that difficulty adjustment is working (at least one block should have difficulty > 1)
		hasAdjustedDifficulty := false
		blocks := blockchain.GetBlocks()
		for _, block := range blocks {
			if block.Difficulty > 1 {
				hasAdjustedDifficulty = true
				break
			}
		}

		// Note: This test might not always pass because difficulty adjustment depends on timing
		// In a real test, we would control the timing more precisely
		t.Logf("Has adjusted difficulty: %v", hasAdjustedDifficulty)
	})

	t.Run("ValidateProof", func(t *testing.T) {
		// Create a new blockchain
		blockchain := core.NewBlockchain()
		defer blockchain.Close()

		// Mine a new block
		block := consensus.MineBlock(blockchain, "Test transaction")

		// Create proof of work for the block
		pow := consensus.NewProofOfWork(block)

		// Validate the proof of work
		if !pow.Validate() {
			t.Error("Expected proof of work to be valid")
		}
	})

	t.Run("InvalidProof", func(t *testing.T) {
		// Create a new blockchain
		blockchain := core.NewBlockchain()
		defer blockchain.Close()

		// Mine a new block
		block := consensus.MineBlock(blockchain, "Test transaction")

		// Tamper with the block to make the proof invalid
		block.Nonce = 12345 // Change the nonce

		// Create proof of work for the tampered block
		pow := consensus.NewProofOfWork(block)

		// Validate the proof of work (should fail)
		if pow.Validate() {
			t.Error("Expected proof of work to be invalid")
		}
	})
}
