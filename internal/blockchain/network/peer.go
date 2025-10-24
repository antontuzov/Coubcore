package network

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// Peer represents a connected node in the P2P network
type Peer struct {
	conn         net.Conn
	addr         string
	messageChan  chan<- Message
	deadPeerChan chan<- *Peer
}

// Message represents a message exchanged between peers
type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// Handshake represents a handshake message between peers
type Handshake struct {
	Version    int    `json:"version"`
	AddrFrom   string `json:"addr_from"`
	AddrTo     string `json:"addr_to"`
	ListenPort int    `json:"listen_port"`
}

// NewPeer creates a new peer
func NewPeer(conn net.Conn, messageChan chan<- Message, deadPeerChan chan<- *Peer) *Peer {
	return &Peer{
		conn:         conn,
		addr:         conn.RemoteAddr().String(),
		messageChan:  messageChan,
		deadPeerChan: deadPeerChan,
	}
}

// Addr returns the address of the peer
func (p *Peer) Addr() string {
	return p.addr
}

// SendHandshake sends a handshake message to the peer
func (p *Peer) SendHandshake(handshake Handshake) error {
	msg := Message{
		Type:    "handshake",
		Payload: handshake,
	}
	return p.SendMessage(msg)
}

// ReceiveHandshake receives a handshake message from the peer
func (p *Peer) ReceiveHandshake() (Handshake, error) {
	var handshake Handshake

	// Read the message length
	lenBuf := make([]byte, 4)
	_, err := io.ReadFull(p.conn, lenBuf)
	if err != nil {
		return handshake, err
	}

	// Convert length to int
	length := int(lenBuf[0])<<24 | int(lenBuf[1])<<16 | int(lenBuf[2])<<8 | int(lenBuf[3])

	// Read the message data
	data := make([]byte, length)
	_, err = io.ReadFull(p.conn, data)
	if err != nil {
		return handshake, err
	}

	// Unmarshal the message
	var msg Message
	err = json.Unmarshal(data, &msg)
	if err != nil {
		return handshake, err
	}

	// Check if it's a handshake message
	if msg.Type != "handshake" {
		return handshake, fmt.Errorf("expected handshake message, got %s", msg.Type)
	}

	// Marshal the payload to JSON and then unmarshal to Handshake
	payloadData, err := json.Marshal(msg.Payload)
	if err != nil {
		return handshake, err
	}

	err = json.Unmarshal(payloadData, &handshake)
	return handshake, err
}

// SendMessage sends a message to the peer
func (p *Peer) SendMessage(msg Message) error {
	// Marshal the message to JSON
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// Prepend the message length
	length := len(data)
	lenBuf := make([]byte, 4)
	lenBuf[0] = byte(length >> 24)
	lenBuf[1] = byte(length >> 16)
	lenBuf[2] = byte(length >> 8)
	lenBuf[3] = byte(length)

	// Send the message length followed by the message data
	_, err = p.conn.Write(append(lenBuf, data...))
	return err
}

// ListenForMessages listens for incoming messages from the peer
func (p *Peer) ListenForMessages() {
	defer func() {
		p.conn.Close()
		p.deadPeerChan <- p
	}()

	for {
		// Set a read deadline to prevent hanging
		p.conn.SetReadDeadline(time.Now().Add(30 * time.Second))

		// Read the message length
		lenBuf := make([]byte, 4)
		_, err := io.ReadFull(p.conn, lenBuf)
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading message length from peer %s: %v", p.addr, err)
			}
			return
		}

		// Convert length to int
		length := int(lenBuf[0])<<24 | int(lenBuf[1])<<16 | int(lenBuf[2])<<8 | int(lenBuf[3])

		// Read the message data
		data := make([]byte, length)
		_, err = io.ReadFull(p.conn, data)
		if err != nil {
			log.Printf("Error reading message data from peer %s: %v", p.addr, err)
			return
		}

		// Unmarshal the message
		var msg Message
		err = json.Unmarshal(data, &msg)
		if err != nil {
			log.Printf("Error unmarshaling message from peer %s: %v", p.addr, err)
			continue
		}

		// Send the message to the message channel
		p.messageChan <- msg
	}
}

// Close closes the peer connection
func (p *Peer) Close() error {
	return p.conn.Close()
}
