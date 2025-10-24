package consensus

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/antontuzov/coubcore/internal/blockchain/core"
)

// ProofOfWork represents a proof of work consensus mechanism
type ProofOfWork struct {
	Block  *core.Block
	Target *big.Int
}

// NewProofOfWork creates a new ProofOfWork
func NewProofOfWork(block *core.Block) *ProofOfWork {
	// Limit difficulty to prevent overflow
	difficulty := block.Difficulty
	if difficulty > 128 {
		difficulty = 128
	}

	target := big.NewInt(1)
	target.Lsh(target, uint(256-int(difficulty)))

	pow := &ProofOfWork{block, target}

	return pow
}

// InitData initializes the data for proof of work
func (pow *ProofOfWork) InitData(nonce uint64) []byte {
	data := bytes.Join(
		[][]byte{
			[]byte(pow.Block.PreviousHash),
			[]byte(fmt.Sprintf("%v", pow.Block.Data)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Difficulty)),
		},
		[]byte{},
	)

	return data
}

// Run performs the proof of work
func (pow *ProofOfWork) Run() (uint64, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := uint64(0)

	fmt.Printf("Mining a new block with difficulty %d\n", pow.Block.Difficulty)
	for nonce < math.MaxUint64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// Validate validates the proof of work
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.Target) == -1
}

// IntToHex converts an int64 to a hex byte slice
func IntToHex(num int64) []byte {
	return []byte(fmt.Sprintf("%x", num))
}

// CalculateDifficulty calculates the new difficulty based on the time it took to mine the last blocks
func CalculateDifficulty(lastBlock *core.Block, previousBlock *core.Block) uint8 {
	// Target time between blocks (in seconds)
	targetTime := 10

	// If this is the genesis block or there's no previous block, return default difficulty
	if previousBlock == nil {
		return 1
	}

	// Calculate the time difference between the last two blocks
	timeDiff := lastBlock.Timestamp.Sub(previousBlock.Timestamp).Seconds()

	// Calculate the new difficulty
	// If blocks are being mined too quickly, increase difficulty
	// If blocks are being mined too slowly, decrease difficulty
	newDifficulty := lastBlock.Difficulty

	if timeDiff < float64(targetTime) {
		// Blocks are being mined too quickly, increase difficulty
		if newDifficulty < 128 {
			newDifficulty++
		}
	} else if timeDiff > float64(targetTime) {
		// Blocks are being mined too slowly, decrease difficulty (but not below 1)
		if newDifficulty > 1 {
			newDifficulty--
		}
	}

	return newDifficulty
}

// MineBlock mines a new block using proof of work
func MineBlock(blockchain *core.Blockchain, data interface{}) *core.Block {
	// Get the previous block
	previousBlock := blockchain.GetLatestBlock()

	// Calculate the new difficulty
	var newDifficulty uint8 = 1
	if previousBlock.Index > 0 {
		prevPrevBlock := blockchain.GetBlockByIndex(previousBlock.Index - 1)
		newDifficulty = CalculateDifficulty(previousBlock, prevPrevBlock)
	}

	// Create a new block with the calculated difficulty
	newBlock := core.NewBlock(previousBlock.Index+1, previousBlock.Hash, data)
	newBlock.Difficulty = newDifficulty
	newBlock.Timestamp = time.Now()

	// Create proof of work for the new block
	pow := NewProofOfWork(newBlock)

	// Run the proof of work
	nonce, hash := pow.Run()

	// Set the nonce and hash for the block
	newBlock.Nonce = nonce
	newBlock.Hash = fmt.Sprintf("%x", hash)

	// Validate the proof of work
	if !pow.Validate() {
		return nil
	}

	// Add the block to the blockchain
	blockchain.AddBlockManually(newBlock)

	return newBlock
}
