package main

import (
	"os"
	"testing"

	"github.com/antontuzov/coubcore/internal/blockchain/consensus"
	"github.com/antontuzov/coubcore/internal/blockchain/core"
)

func TestBlockchainConsensus(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	t.Run("ValidChain", func(t *testing.T) {
		// Create a new blockchain
		blockchain := core.NewBlockchain()
		defer blockchain.Close()

		// Add some blocks
		blockchain.AddBlock("Block 1")
		blockchain.AddBlock("Block 2")
		blockchain.AddBlock("Block 3")

		// Validate the chain
		if !blockchain.IsChainValid() {
			t.Error("Expected blockchain to be valid")
		}
	})

	t.Run("InvalidChain", func(t *testing.T) {
		// Create a new blockchain
		blockchain := core.NewBlockchain()
		defer blockchain.Close()

		// Add some blocks
		blockchain.AddBlock("Block 1")
		blockchain.AddBlock("Block 2")

		// Tamper with a block to make it invalid
		block := blockchain.GetBlockByIndex(1)
		block.Data = "Tampered Data"

		// Validate the chain (should fail)
		if blockchain.IsChainValid() {
			t.Error("Expected blockchain to be invalid after tampering")
		}
	})

	t.Run("ProofOfWorkValidation", func(t *testing.T) {
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

	t.Run("ChainReplacement", func(t *testing.T) {
		// Create two blockchains
		bc1 := core.NewBlockchain()
		bc2 := core.NewBlockchain()
		defer bc1.Close()
		defer bc2.Close()

		// Add more blocks to bc2
		bc2.AddBlock("Block 1")
		bc2.AddBlock("Block 2")

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
	})

	t.Run("ChainReplacementShorterChain", func(t *testing.T) {
		// Create two blockchains
		bc1 := core.NewBlockchain()
		bc2 := core.NewBlockchain()
		defer bc1.Close()
		defer bc2.Close()

		// Add more blocks to bc1
		bc1.AddBlock("Block 1")
		bc1.AddBlock("Block 2")

		// Get blocks from bc2 (shorter chain)
		blocks := bc2.GetBlocks()

		// Try to replace bc1's chain with bc2's chain (should fail)
		success := bc1.ReplaceChain(blocks)

		// Check that the replacement failed
		if success {
			t.Error("Expected chain replacement to fail with shorter chain")
		}
	})
}
