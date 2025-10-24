package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

// Wallet represents a cryptocurrency wallet
type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  []byte
	Address    string
}

// NewWallet creates a new wallet with a new key pair
func NewWallet() (*Wallet, error) {
	// Generate a new private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	// Get the public key
	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	// Generate the address
	address := generateAddress(publicKey)

	wallet := &Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    address,
	}

	return wallet, nil
}

// generateAddress generates a wallet address from a public key
func generateAddress(publicKey []byte) string {
	// SHA256 hash of the public key
	pubKeyHash := sha256.Sum256(publicKey)

	// Simple address generation (in a real implementation, we would use RIPEMD160)
	addressBytes := pubKeyHash[:20] // Take first 20 bytes

	// Convert to hex string for simplicity
	address := hex.EncodeToString(addressBytes)

	return address
}

// ValidateAddress validates a wallet address
func ValidateAddress(address string) bool {
	// Simple validation - check if it's a valid hex string of appropriate length
	if len(address) != 40 {
		return false
	}

	// Try to decode
	_, err := hex.DecodeString(address)
	return err == nil
}

// Sign signs data with the wallet's private key
func (w *Wallet) Sign(data []byte) ([]byte, error) {
	// Hash the data
	hash := sha256.Sum256(data)

	// Sign the hash
	r, s, err := ecdsa.Sign(rand.Reader, w.PrivateKey, hash[:])
	if err != nil {
		return nil, err
	}

	// Concatenate r and s
	signature := append(r.Bytes(), s.Bytes()...)

	return signature, nil
}

// Verify verifies a signature with the wallet's public key
func (w *Wallet) Verify(data []byte, signature []byte) bool {
	// Hash the data
	hash := sha256.Sum256(data)

	// Split signature into r and s
	r := new(big.Int).SetBytes(signature[:len(signature)/2])
	s := new(big.Int).SetBytes(signature[len(signature)/2:])

	// Verify the signature
	return ecdsa.Verify(&w.PrivateKey.PublicKey, hash[:], r, s)
}

// GetBalance calculates the balance of the wallet (simplified version)
func (w *Wallet) GetBalance() int {
	// In a real implementation, this would calculate balance from UTXO set
	// For now, we'll return 0
	return 0
}
