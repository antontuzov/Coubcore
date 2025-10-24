package main

import (
	"fmt"
	"os"

	"github.com/antontuzov/coubcore/internal/blockchain/core"
)

func main() {
	fmt.Println("Testing blockchain...")

	// Create a new blockchain
	blockchain := core.NewBlockchain()
	defer blockchain.Close()

	fmt.Printf("Blockchain length: %d\n", blockchain.Length())

	// Get latest block
	latest := blockchain.GetLatestBlock()
	if latest != nil {
		fmt.Printf("Latest block #%d: %s\n", latest.Index, latest.Hash)
	}

	// Add a block
	fmt.Println("Adding a new block...")
	block := blockchain.AddBlock("Test transaction")
	fmt.Printf("Added block #%d with hash %s\n", block.Index, block.Hash)

	// Check length again
	fmt.Printf("Blockchain length: %d\n", blockchain.Length())

	// Validate chain
	fmt.Println("Validating blockchain...")
	if blockchain.IsChainValid() {
		fmt.Println("Blockchain is valid")
	} else {
		fmt.Println("Blockchain is invalid")
	}

	// List blocks
	blocks := blockchain.GetBlocks()
	fmt.Printf("Listing %d blocks:\n", len(blocks))
	for _, b := range blocks {
		fmt.Printf("Block #%d: %s\n", b.Index, b.Hash)
	}

	os.Exit(0)
}
