package network

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/antontuzov/coubcore/internal/blockchain/core"
)

// Server represents a P2P node in the blockchain network
type Server struct {
	host         string
	port         int
	listener     net.Listener
	blockchain   *core.Blockchain
	peers        map[string]*Peer
	peersMutex   sync.RWMutex
	newPeerChan  chan *Peer
	deadPeerChan chan *Peer
	messageChan  chan Message
}

// NewServer creates a new P2P server
func NewServer(host string, port int, blockchain *core.Blockchain) *Server {
	return &Server{
		host:         host,
		port:         port,
		blockchain:   blockchain,
		peers:        make(map[string]*Peer),
		newPeerChan:  make(chan *Peer),
		deadPeerChan: make(chan *Peer),
		messageChan:  make(chan Message),
	}
}

// Start starts the P2P server
func (s *Server) Start() error {
	// Start listening for incoming connections
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		return err
	}
	s.listener = listener

	log.Printf("P2P server started on %s:%d", s.host, s.port)

	// Start goroutines for handling peers and messages
	go s.handleNewPeers()
	go s.handleDeadPeers()
	go s.handleMessages()

	// Accept incoming connections
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Handle the new connection in a separate goroutine
		go s.handleConnection(conn)
	}
}

// handleConnection handles a new incoming connection
func (s *Server) handleConnection(conn net.Conn) {
	// Create a new peer
	peer := NewPeer(conn, s.messageChan, s.deadPeerChan)

	// Send handshake
	handshake := Handshake{
		Version:    1,
		AddrFrom:   fmt.Sprintf("%s:%d", s.host, s.port),
		AddrTo:     conn.RemoteAddr().String(),
		ListenPort: s.port,
	}

	if err := peer.SendHandshake(handshake); err != nil {
		log.Printf("Error sending handshake: %v", err)
		conn.Close()
		return
	}

	// Wait for handshake response
	response, err := peer.ReceiveHandshake()
	if err != nil {
		log.Printf("Error receiving handshake: %v", err)
		conn.Close()
		return
	}

	log.Printf("Received handshake from %s", response.AddrFrom)

	// Add peer to the server
	s.newPeerChan <- peer
}

// handleNewPeers handles new peer connections
func (s *Server) handleNewPeers() {
	for peer := range s.newPeerChan {
		s.peersMutex.Lock()
		s.peers[peer.Addr()] = peer
		s.peersMutex.Unlock()

		log.Printf("New peer connected: %s", peer.Addr())

		// Send latest block to the new peer
		latestBlock := s.blockchain.GetLatestBlock()
		if latestBlock != nil {
			msg := Message{
				Type:    "latest_block",
				Payload: latestBlock,
			}
			peer.SendMessage(msg)
		}
	}
}

// handleDeadPeers handles dead peer connections
func (s *Server) handleDeadPeers() {
	for peer := range s.deadPeerChan {
		s.peersMutex.Lock()
		delete(s.peers, peer.Addr())
		s.peersMutex.Unlock()

		log.Printf("Peer disconnected: %s", peer.Addr())
	}
}

// handleMessages handles incoming messages from peers
func (s *Server) handleMessages() {
	for msg := range s.messageChan {
		switch msg.Type {
		case "block":
			s.handleBlockMessage(msg)
		case "transaction":
			s.handleTransactionMessage(msg)
		case "get_blocks":
			s.handleGetBlocksMessage(msg)
		case "inventory":
			s.handleInventoryMessage(msg)
		default:
			log.Printf("Unknown message type: %s", msg.Type)
		}
	}
}

// handleBlockMessage handles incoming block messages
func (s *Server) handleBlockMessage(msg Message) {
	// Deserialize the block
	blockData, err := json.Marshal(msg.Payload)
	if err != nil {
		log.Printf("Error marshaling block: %v", err)
		return
	}

	block, err := core.Deserialize(blockData)
	if err != nil {
		log.Printf("Error deserializing block: %v", err)
		return
	}

	// Add the block to the blockchain
	// Note: In a real implementation, we would validate the block before adding it
	s.blockchain.AddBlockManually(block)

	log.Printf("Received and added block #%d", block.Index)
}

// handleTransactionMessage handles incoming transaction messages
func (s *Server) handleTransactionMessage(msg Message) {
	// In a real implementation, we would deserialize and validate the transaction
	log.Printf("Received transaction message")
}

// handleGetBlocksMessage handles get_blocks messages
func (s *Server) handleGetBlocksMessage(msg Message) {
	// In a real implementation, we would send the requested blocks to the peer
	log.Printf("Received get_blocks message")
}

// handleInventoryMessage handles inventory messages
func (s *Server) handleInventoryMessage(msg Message) {
	// In a real implementation, we would request the items in the inventory
	log.Printf("Received inventory message")
}

// ConnectToPeer connects to a remote peer
func (s *Server) ConnectToPeer(address string) error {
	// Connect to the remote peer
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}

	// Create a new peer
	peer := NewPeer(conn, s.messageChan, s.deadPeerChan)

	// Send handshake
	handshake := Handshake{
		Version:    1,
		AddrFrom:   fmt.Sprintf("%s:%d", s.host, s.port),
		AddrTo:     address,
		ListenPort: s.port,
	}

	if err := peer.SendHandshake(handshake); err != nil {
		conn.Close()
		return err
	}

	// Add peer to the server
	s.newPeerChan <- peer

	return nil
}

// BroadcastMessage broadcasts a message to all connected peers
func (s *Server) BroadcastMessage(msg Message) {
	s.peersMutex.RLock()
	defer s.peersMutex.RUnlock()

	for _, peer := range s.peers {
		go func(p *Peer) {
			if err := p.SendMessage(msg); err != nil {
				log.Printf("Error sending message to peer %s: %v", p.Addr(), err)
			}
		}(peer)
	}
}

// Stop stops the P2P server
func (s *Server) Stop() error {
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}
