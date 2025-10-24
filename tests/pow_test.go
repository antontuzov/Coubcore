package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/antontuzov/coubcore/internal/blockchain/consensus"
	"github.com/antontuzov/coubcore/internal/blockchain/core"
)

func TestProofOfWork(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	// Create a new blockchain
	blockchain := core.NewBlockchain()
	defer blockchain.Close()

	// Create a new block
	block := core.NewBlock(1, blockchain.GetLatestBlock().Hash, "Test data")
	block.Difficulty = 1 // Low difficulty for testing

	// Create proof of work
	pow := consensus.NewProofOfWork(block)

	// Run the proof of work
	nonce, hash := pow.Run()

	// Set the nonce and hash
	block.Nonce = nonce
	block.Hash = fmt.Sprintf("%x", hash)

	// Validate the proof of work
	if !pow.Validate() {
		t.Error("Expected proof of work to be valid")
	}

	t.Logf("Nonce: %d", nonce)
	t.Logf("Hash: %x", hash)
}

func TestMineBlock(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	// Create a new blockchain
	blockchain := core.NewBlockchain()
	defer blockchain.Close()

	// Mine a new block
	minedBlock := consensus.MineBlock(blockchain, "Test transaction")

	// Check that the block was mined
	if minedBlock == nil {
		t.Fatal("Expected block to be mined")
	}

	t.Logf("Mined block nonce: %d", minedBlock.Nonce)
	t.Logf("Mined block hash: %s", minedBlock.Hash)

	// Check that the block has a hash
	if minedBlock.Hash == "" {
		t.Error("Expected block to have a hash")
	}

	// Check that the block was added to the blockchain
	if blockchain.Length() != 2 {
		t.Errorf("Expected blockchain length to be 2, got %d", blockchain.Length())
	}

	// Get the latest block from the blockchain
	latestBlock := blockchain.GetLatestBlock()
	if latestBlock.Hash != minedBlock.Hash {
		t.Error("Expected latest block to match mined block")
	}

	t.Logf("Mined block #%d with hash %s", minedBlock.Index, minedBlock.Hash)
}
