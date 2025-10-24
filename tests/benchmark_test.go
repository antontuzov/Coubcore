package main

import (
	"os"
	"testing"

	"github.com/antontuzov/coubcore/internal/blockchain/consensus"
	"github.com/antontuzov/coubcore/internal/blockchain/core"
	"github.com/antontuzov/coubcore/internal/blockchain/wallet"
)

func BenchmarkBlockMining(b *testing.B) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	// Create a new blockchain
	blockchain := core.NewBlockchain()
	defer blockchain.Close()

	// Reset the benchmark timer
	b.ResetTimer()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Mine a new block
		block := consensus.MineBlock(blockchain, "Benchmark transaction")
		if block == nil {
			b.Fatal("Failed to mine block")
		}
	}
}

func BenchmarkBlockValidation(b *testing.B) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	// Create a new blockchain
	blockchain := core.NewBlockchain()
	defer blockchain.Close()

	// Mine a block for testing
	block := consensus.MineBlock(blockchain, "Test transaction")
	if block == nil {
		b.Fatal("Failed to mine block")
	}

	// Reset the benchmark timer
	b.ResetTimer()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Validate the block
		if !blockchain.IsChainValid() {
			b.Fatal("Block validation failed")
		}
	}
}

func BenchmarkTransactionCreation(b *testing.B) {
	// Create a UTXO set
	utxoSet := make(map[string][]core.UTXO)

	// Reset the benchmark timer
	b.ResetTimer()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Create a transaction
		_, err := core.NewTransaction("sender", "recipient", 100, utxoSet)
		if err != nil {
			b.Fatal("Failed to create transaction")
		}
	}
}

func BenchmarkWalletCreation(b *testing.B) {
	// Reset the benchmark timer
	b.ResetTimer()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Create a wallet
		_, err := wallet.NewWallet()
		if err != nil {
			b.Fatal("Failed to create wallet")
		}
	}
}

func BenchmarkAddressValidation(b *testing.B) {
	// Create a wallet to get a valid address
	w, err := wallet.NewWallet()
	if err != nil {
		b.Fatal("Failed to create wallet")
	}

	// Reset the benchmark timer
	b.ResetTimer()

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Validate the address
		wallet.ValidateAddress(w.Address)
	}
}
