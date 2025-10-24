package core

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Index        uint64      `json:"index"`
	Timestamp    time.Time   `json:"timestamp"`
	PreviousHash string      `json:"previousHash"`
	Hash         string      `json:"hash"`
	Data         interface{} `json:"data"`
	Nonce        uint64      `json:"nonce"`
	Difficulty   uint8       `json:"difficulty"`
	Validator    string      `json:"validator"`
}

// NewBlock creates a new block
func NewBlock(index uint64, previousHash string, data interface{}) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now(),
		PreviousHash: previousHash,
		Data:         data,
		Nonce:        0,
		Difficulty:   1,
		Validator:    "",
	}

	// Calculate the hash for the new block
	block.Hash = block.CalculateHash()

	return block
}

// CalculateHash calculates the SHA256 hash of the block
func (b *Block) CalculateHash() string {
	// Create a copy of the block without the hash field to avoid circular dependency
	blockCopy := Block{
		Index:        b.Index,
		Timestamp:    b.Timestamp,
		PreviousHash: b.PreviousHash,
		Data:         b.Data,
		Nonce:        b.Nonce,
		Difficulty:   b.Difficulty,
		Validator:    b.Validator,
		Hash:         "", // Set to empty to avoid circular dependency
	}

	// Convert the block copy to JSON for hashing
	blockData, err := json.Marshal(blockCopy)
	if err != nil {
		return ""
	}

	// Calculate SHA256 hash
	hash := sha256.Sum256(blockData)
	return fmt.Sprintf("%x", hash)
}

// Serialize converts the block to a JSON byte array
func (b *Block) Serialize() ([]byte, error) {
	return json.Marshal(b)
}

// Deserialize converts a JSON byte array to a block
func Deserialize(data []byte) (*Block, error) {
	var block Block
	err := json.Unmarshal(data, &block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

// Validate checks the integrity of the block
func (b *Block) Validate(previousBlock *Block) bool {
	// Check if the hash is correct
	if b.CalculateHash() != b.Hash {
		return false
	}

	// Check if the previous hash matches the previous block's hash
	if previousBlock != nil && b.PreviousHash != previousBlock.Hash {
		return false
	}

	// Check if the index is correct
	if previousBlock != nil && b.Index != previousBlock.Index+1 {
		return false
	}

	// Check if the timestamp is after the previous block's timestamp
	if previousBlock != nil && !b.Timestamp.After(previousBlock.Timestamp) {
		return false
	}

	return true
}
