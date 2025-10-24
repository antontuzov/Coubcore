package core

import (
	"fmt"
	"sync"

	"go.etcd.io/bbolt"
)

// Blockchain represents the entire blockchain
type Blockchain struct {
	blocks []*Block
	db     *bbolt.DB
	mu     sync.RWMutex
}

// NewBlockchain creates a new blockchain with a genesis block
func NewBlockchain() *Blockchain {
	// Open the database
	db, err := bbolt.Open("blockchain.db", 0600, nil)
	if err != nil {
		panic(err)
	}

	bc := &Blockchain{
		blocks: make([]*Block, 0),
		db:     db,
	}

	// Initialize the database buckets
	bc.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("blocks"))
		return err
	})

	// Load existing blockchain from database or create genesis block
	bc.loadFromDB()

	// If blockchain is empty, create genesis block
	if len(bc.blocks) == 0 {
		genesisBlock := NewBlock(0, "", "Genesis Block")
		bc.blocks = append(bc.blocks, genesisBlock)
		bc.persistBlock(genesisBlock)
	}

	return bc
}

// loadFromDB loads the blockchain from the database
func (bc *Blockchain) loadFromDB() {
	bc.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("blocks"))
		if bucket == nil {
			return nil
		}

		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			block, err := Deserialize(v)
			if err != nil {
				continue
			}
			bc.blocks = append(bc.blocks, block)
		}

		return nil
	})
}

// persistBlock saves a block to the database
func (bc *Blockchain) persistBlock(block *Block) error {
	return bc.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("blocks"))
		if bucket == nil {
			return fmt.Errorf("blocks bucket not found")
		}

		data, err := block.Serialize()
		if err != nil {
			return err
		}

		key := fmt.Sprintf("%020d", block.Index)
		return bucket.Put([]byte(key), data)
	})
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data interface{}) *Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	// Get the previous block
	previousBlock := bc.blocks[len(bc.blocks)-1]

	// Create a new block
	newBlock := NewBlock(previousBlock.Index+1, previousBlock.Hash, data)

	// Add the new block to the chain
	bc.blocks = append(bc.blocks, newBlock)

	// Persist the new block
	bc.persistBlock(newBlock)

	return newBlock
}

// AddBlockManually adds a pre-created block to the blockchain
func (bc *Blockchain) AddBlockManually(block *Block) {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	// Add the block to the chain
	bc.blocks = append(bc.blocks, block)

	// Persist the new block
	bc.persistBlock(block)
}

// GetLatestBlock returns the latest block in the blockchain
func (bc *Blockchain) GetLatestBlock() *Block {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	if len(bc.blocks) == 0 {
		return nil
	}

	return bc.blocks[len(bc.blocks)-1]
}

// GetBlockByIndex returns a block by its index
func (bc *Blockchain) GetBlockByIndex(index uint64) *Block {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	if index >= uint64(len(bc.blocks)) {
		return nil
	}

	return bc.blocks[index]
}

// IsChainValid validates the entire blockchain
func (bc *Blockchain) IsChainValid() bool {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	// Check if the blockchain is empty
	if len(bc.blocks) == 0 {
		return false
	}

	// Validate each block in the chain
	for i := 1; i < len(bc.blocks); i++ {
		currentBlock := bc.blocks[i]
		previousBlock := bc.blocks[i-1]

		// Validate the current block
		if !currentBlock.Validate(previousBlock) {
			fmt.Printf("Invalid block at index %d\n", currentBlock.Index)
			return false
		}
	}

	return true
}

// GetBlocks returns all blocks in the blockchain
func (bc *Blockchain) GetBlocks() []*Block {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	// Return a copy of the blocks slice
	blocks := make([]*Block, len(bc.blocks))
	copy(blocks, bc.blocks)

	return blocks
}

// Length returns the number of blocks in the blockchain
func (bc *Blockchain) Length() int {
	bc.mu.RLock()
	defer bc.mu.RUnlock()

	return len(bc.blocks)
}

// ReplaceChain replaces the current chain with a longer valid chain
func (bc *Blockchain) ReplaceChain(newChain []*Block) bool {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	// Check if the new chain is longer than the current chain
	if len(newChain) <= len(bc.blocks) {
		return false
	}

	// Validate the new chain
	// Create a temporary blockchain with the new chain
	tempBC := &Blockchain{
		blocks: newChain,
		db:     bc.db, // Use the same database connection
	}

	if !tempBC.IsChainValid() {
		return false
	}

	// Replace the current chain with the new chain
	bc.blocks = newChain

	// Persist the new chain
	bc.db.Update(func(tx *bbolt.Tx) error {
		// Clear the blocks bucket
		tx.DeleteBucket([]byte("blocks"))
		bucket, err := tx.CreateBucket([]byte("blocks"))
		if err != nil {
			return err
		}

		// Add all blocks to the bucket
		for _, block := range newChain {
			data, err := block.Serialize()
			if err != nil {
				continue
			}

			key := fmt.Sprintf("%020d", block.Index)
			err = bucket.Put([]byte(key), data)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return true
}

// Close closes the database connection
func (bc *Blockchain) Close() error {
	return bc.db.Close()
}
