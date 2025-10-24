package main

import (
	"testing"

	"github.com/antontuzov/coubcore/internal/blockchain/core"
)

func TestTransactionCreation(t *testing.T) {
	// Create a coinbase transaction
	cbTx := core.NewCoinbaseTransaction("recipient", "test coinbase")

	// Check that the transaction was created
	if cbTx == nil {
		t.Fatal("Expected coinbase transaction to be created")
	}

	// Check that the transaction has an ID
	if cbTx.ID == "" {
		t.Error("Expected transaction to have an ID")
	}

	// Check that the transaction is a coinbase transaction
	if !cbTx.IsCoinbase() {
		t.Error("Expected transaction to be a coinbase transaction")
	}

	// Check that the transaction has inputs
	if len(cbTx.Inputs) == 0 {
		t.Error("Expected transaction to have inputs")
	}

	// Check that the transaction has outputs
	if len(cbTx.Outputs) == 0 {
		t.Error("Expected transaction to have outputs")
	}

	t.Logf("Coinbase transaction ID: %s", cbTx.ID)
}

func TestRegularTransactionCreation(t *testing.T) {
	// Create a regular transaction
	utxoSet := make(map[string][]core.UTXO)
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

	t.Logf("Regular transaction ID: %s", tx.ID)
}

func TestTransactionHashing(t *testing.T) {
	// Create a coinbase transaction
	cbTx := core.NewCoinbaseTransaction("recipient", "test coinbase")

	// Get the hash
	hash := cbTx.Hash()

	// Check that we got a hash
	if len(hash) == 0 {
		t.Error("Expected hash to exist")
	}

	t.Logf("Transaction hash: %x", hash)
}

func TestTrimmedCopy(t *testing.T) {
	// Create a coinbase transaction
	cbTx := core.NewCoinbaseTransaction("recipient", "test coinbase")

	// Create a trimmed copy
	trimmed := cbTx.TrimmedCopy()

	// Check that the trimmed copy has the same ID
	if trimmed.ID != cbTx.ID {
		t.Error("Expected trimmed copy to have the same ID")
	}

	// Check that the trimmed copy has the same number of inputs and outputs
	if len(trimmed.Inputs) != len(cbTx.Inputs) {
		t.Error("Expected trimmed copy to have the same number of inputs")
	}

	if len(trimmed.Outputs) != len(cbTx.Outputs) {
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
}
