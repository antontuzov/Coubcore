package core

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

// Transaction represents a transaction in the blockchain
type Transaction struct {
	ID      string     `json:"id"`
	Inputs  []TXInput  `json:"inputs"`
	Outputs []TXOutput `json:"outputs"`
	Time    time.Time  `json:"time"`
}

// TXInput represents a transaction input
type TXInput struct {
	TXID      string `json:"txid"`
	Vout      int    `json:"vout"`
	Signature []byte `json:"signature"`
	PubKey    []byte `json:"pubKey"`
}

// TXOutput represents a transaction output
type TXOutput struct {
	Value      int    `json:"value"`
	PubKeyHash []byte `json:"pubKeyHash"`
}

// UTXO represents an unspent transaction output
type UTXO struct {
	TXID   string
	Index  int
	Output TXOutput
}

// NewTransaction creates a new transaction
func NewTransaction(from, to string, amount int, utxoSet map[string][]UTXO) (*Transaction, error) {
	// Create a regular transaction with proper inputs and outputs
	txin := TXInput{
		TXID:      "previous_tx_id", // In a real implementation, this would be a real TXID
		Vout:      0,                // In a real implementation, this would be the output index
		Signature: nil,
		PubKey:    []byte(from),
	}

	txout := TXOutput{
		Value:      amount,
		PubKeyHash: []byte(to),
	}

	tx := Transaction{
		Inputs:  []TXInput{txin},
		Outputs: []TXOutput{txout},
		Time:    time.Now(),
	}

	tx.SetID()

	return &tx, nil
}

// NewCoinbaseTransaction creates a coinbase transaction
func NewCoinbaseTransaction(to, data string) *Transaction {
	if data == "" {
		data = "Coinbase Transaction"
	}

	txin := TXInput{
		TXID:      "",
		Vout:      -1,
		Signature: []byte(data),
		PubKey:    []byte(data),
	}

	txout := TXOutput{
		Value:      100, // Mining reward
		PubKeyHash: []byte(to),
	}

	tx := Transaction{
		Inputs:  []TXInput{txin},
		Outputs: []TXOutput{txout},
		Time:    time.Now(),
	}

	tx.SetID()

	return &tx
}

// SetID sets the ID of the transaction
func (tx *Transaction) SetID() {
	txBytes, err := json.Marshal(tx)
	if err != nil {
		return
	}

	hash := sha256.Sum256(txBytes)
	tx.ID = string(hash[:])
}

// IsCoinbase checks if the transaction is a coinbase transaction
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].TXID) == 0 && tx.Inputs[0].Vout == -1
}

// Sign signs each input of the transaction
func (tx *Transaction) Sign(privKey interface{}, prevTXs map[string]Transaction) error {
	if tx.IsCoinbase() {
		return nil
	}

	// TODO: Implement transaction signing
	// This would involve:
	// 1. Creating a copy of the transaction
	// 2. Signing each input with the private key
	// 3. Setting the signature in each input

	return nil
}

// Verify verifies the signatures of the transaction
func (tx *Transaction) Verify(prevTXs map[string]Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}

	// TODO: Implement transaction verification
	// This would involve:
	// 1. Verifying each input's signature
	// 2. Checking that the signatures match the public keys

	return true
}

// Hash returns the hash of the transaction
func (tx *Transaction) Hash() []byte {
	txCopy := *tx
	txCopy.ID = ""

	txBytes, err := json.Marshal(txCopy)
	if err != nil {
		return nil
	}

	hash := sha256.Sum256(txBytes)
	return hash[:]
}

// TrimmedCopy returns a copy of the transaction with empty signatures and pubkeys
func (tx *Transaction) TrimmedCopy() Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	for _, vin := range tx.Inputs {
		inputs = append(inputs, TXInput{vin.TXID, vin.Vout, nil, nil})
	}

	for _, vout := range tx.Outputs {
		outputs = append(outputs, TXOutput{vout.Value, vout.PubKeyHash})
	}

	txCopy := Transaction{tx.ID, inputs, outputs, tx.Time}
	return txCopy
}
