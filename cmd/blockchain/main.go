package main

import (
	"fmt"
	"os"

	"github.com/antontuzov/coubcore/internal/blockchain/api"
	"github.com/antontuzov/coubcore/internal/blockchain/core"
	"github.com/antontuzov/coubcore/internal/blockchain/network"
)

func main() {
	fmt.Println("Coubcore Blockchain Node")
	fmt.Println("========================")

	// Create a new blockchain
	blockchain := core.NewBlockchain()
	defer blockchain.Close()

	// Create a new network server
	networkServer := network.NewServer("localhost", 8000, blockchain)

	// Create a new API server
	apiServer := api.NewServer(blockchain, networkServer, 8080)

	// Start the network server in a separate goroutine
	go func() {
		if err := networkServer.Start(); err != nil {
			fmt.Printf("Error starting network server: %v\n", err)
		}
	}()

	// Start the API server in a separate goroutine
	go func() {
		if err := apiServer.Start(); err != nil {
			fmt.Printf("Error starting API server: %v\n", err)
		}
	}()

	// If we have command line arguments, process them
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "add":
			if len(os.Args) > 2 {
				data := os.Args[2]
				block := blockchain.AddBlock(data)
				fmt.Printf("Added block #%d with hash %s\n", block.Index, block.Hash)
			} else {
				fmt.Println("Usage: blockchain add <data>")
			}
		case "list":
			blocks := blockchain.GetBlocks()
			fmt.Printf("Blockchain length: %d\n", len(blocks))
			for _, block := range blocks {
				fmt.Printf("Block #%d: %s\n", block.Index, block.Hash)
			}
		case "validate":
			if blockchain.IsChainValid() {
				fmt.Println("Blockchain is valid")
			} else {
				fmt.Println("Blockchain is invalid")
			}
		case "connect":
			if len(os.Args) > 2 {
				address := os.Args[2]
				if err := networkServer.ConnectToPeer(address); err != nil {
					fmt.Printf("Error connecting to peer %s: %v\n", address, err)
				} else {
					fmt.Printf("Connected to peer %s\n", address)
				}
			} else {
				fmt.Println("Usage: blockchain connect <address>")
			}
		default:
			fmt.Printf("Unknown command: %s\n", os.Args[1])
			fmt.Println("Available commands: add, list, validate, connect")
		}
	} else {
		// Print basic info
		fmt.Printf("Blockchain length: %d\n", blockchain.Length())
		if latest := blockchain.GetLatestBlock(); latest != nil {
			fmt.Printf("Latest block hash: %s\n", latest.Hash)
		}
		fmt.Println("Available commands: add, list, validate, connect")
		fmt.Println("Network server running on port 8000")
		fmt.Println("API server running on port 8080")
	}

	// Keep the program running
	select {}
}
