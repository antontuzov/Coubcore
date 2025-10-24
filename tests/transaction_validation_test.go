package main

import (
	"os"
	"testing"

	"github.com/antontuzov/coubcore/internal/blockchain/core"
)

func TestTransactionValidation(t *testing.T) {
	// Remove existing database for clean test
	os.Remove("blockchain.db")

	t.Run("CoinbaseTransaction", func(t *testing.T) {
		// Create a coinbase transaction
		tx := core.NewCoinbaseTransaction("recipient", "test coinbase")

		// Check that the transaction was created
		if tx == nil {
			t.Fatal("Expected coinbase transaction to be created")
		}

		// Check that the transaction is a coinbase transaction
		if !tx.IsCoinbase() {
			t.Error("Expected transaction to be a coinbase transaction")
		}

		// Check that the transaction has an ID
		if tx.ID == "" {
			t.Error("Expected transaction to have an ID")
		}
	})

	t.Run("RegularTransaction", func(t *testing.T) {
		// Create a UTXO set
		utxoSet := make(map[string][]core.UTXO)

		// Create a regular transaction
		tx, err := core.NewTransaction("sender", "recipient", 100, utxoSet)

		// Check that the transaction was created without error
		if err != nil {
			t.Fatalf("Failed to create transaction: %v", err)
		}

		if tx == nil {
			t.Fatal("Expected transaction to be created")
		}

		// Check that the transaction has an ID
		if tx.ID == "" {
			t.Error("Expected transaction to have an ID")
		}

		// Check that the transaction is not a coinbase transaction
		if tx.IsCoinbase() {
			t.Error("Expected transaction to not be a coinbase transaction")
		}
	})

	t.Run("TransactionHashing", func(t *testing.T) {
		// Create a coinbase transaction
		tx := core.NewCoinbaseTransaction("recipient", "test coinbase")

		// Get the hash
		hash := tx.Hash()

		// Check that we got a hash
		if len(hash) == 0 {
			t.Error("Expected hash to exist")
		}

		// Check that the hash is consistent
		hash2 := tx.Hash()
		if string(hash) != string(hash2) {
			t.Error("Expected hash to be consistent")
		}
	})

	t.Run("TrimmedCopy", func(t *testing.T) {
		// Create a coinbase transaction
		tx := core.NewCoinbaseTransaction("recipient", "test coinbase")

		// Create a trimmed copy
		trimmed := tx.TrimmedCopy()

		// Check that the trimmed copy has the same ID
		if trimmed.ID != tx.ID {
			t.Error("Expected trimmed copy to have the same ID")
		}

		// Check that the trimmed copy has the same number of inputs and outputs
		if len(trimmed.Inputs) != len(tx.Inputs) {
			t.Error("Expected trimmed copy to have the same number of inputs")
		}

		if len(trimmed.Outputs) != len(tx.Outputs) {
			t.Error("Expected trimmed copy to have the same number of outputs")
		}

		// Check that signatures and pubkeys are nil in the trimmed copy
		for _, input := range trimmed.Inputs {
			if input.Signature != nil {
				t.Error("Expected signature to be nil in trimmed copy")
			}
			if input.PubKey != nil {
				t.Error("Expected pubkey to be nil in trimmed copy")
			}
		}
	})
}
