package main

import (
	"testing"

	"github.com/antontuzov/coubcore/internal/blockchain/wallet"
)

func TestWalletCryptography(t *testing.T) {
	t.Run("WalletCreation", func(t *testing.T) {
		// Create a new wallet
		w, err := wallet.NewWallet()
		if err != nil {
			t.Fatalf("Failed to create wallet: %v", err)
		}

		// Check that the wallet was created
		if w == nil {
			t.Fatal("Expected wallet to be created")
		}

		// Check that the private key exists
		if w.PrivateKey == nil {
			t.Error("Expected private key to exist")
		}

		// Check that the public key exists
		if len(w.PublicKey) == 0 {
			t.Error("Expected public key to exist")
		}

		// Check that the address exists
		if w.Address == "" {
			t.Error("Expected address to exist")
		}
	})

	t.Run("AddressValidation", func(t *testing.T) {
		// Create a new wallet
		w, err := wallet.NewWallet()
		if err != nil {
			t.Fatalf("Failed to create wallet: %v", err)
		}

		// Validate the wallet's address
		if !wallet.ValidateAddress(w.Address) {
			t.Error("Expected wallet address to be valid")
		}

		// Test invalid address
		invalidAddress := "invalid_address"
		if wallet.ValidateAddress(invalidAddress) {
			t.Error("Expected invalid address to be rejected")
		}

		// Test address with wrong length
		wrongLengthAddress := "abc123"
		if wallet.ValidateAddress(wrongLengthAddress) {
			t.Error("Expected wrong length address to be rejected")
		}
	})

	t.Run("SigningAndVerification", func(t *testing.T) {
		// Create a new wallet
		w, err := wallet.NewWallet()
		if err != nil {
			t.Fatalf("Failed to create wallet: %v", err)
		}

		// Data to sign
		data := []byte("Hello, blockchain!")

		// Sign the data
		signature, err := w.Sign(data)
		if err != nil {
			t.Fatalf("Failed to sign data: %v", err)
		}

		// Check that we got a signature
		if len(signature) == 0 {
			t.Error("Expected signature to exist")
		}

		// Verify the signature
		if !w.Verify(data, signature) {
			t.Error("Expected signature to be valid")
		}

		// Test with different data (should fail)
		differentData := []byte("Different data")
		if w.Verify(differentData, signature) {
			t.Error("Expected signature to be invalid for different data")
		}
	})

	t.Run("InvalidSignature", func(t *testing.T) {
		// Create a new wallet
		w, err := wallet.NewWallet()
		if err != nil {
			t.Fatalf("Failed to create wallet: %v", err)
		}

		// Data to sign
		data := []byte("Hello, blockchain!")

		// Create an invalid signature
		invalidSignature := []byte("invalid signature")

		// Verify the invalid signature (should fail)
		if w.Verify(data, invalidSignature) {
			t.Error("Expected invalid signature to be rejected")
		}
	})

	t.Run("GetBalance", func(t *testing.T) {
		// Create a new wallet
		w, err := wallet.NewWallet()
		if err != nil {
			t.Fatalf("Failed to create wallet: %v", err)
		}

		// Get the balance (should be 0 for now)
		balance := w.GetBalance()
		if balance != 0 {
			t.Errorf("Expected balance to be 0, got %d", balance)
		}
	})
}
