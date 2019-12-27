package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type Node struct {
	sync.Mutex
	chain   *Chain
	peers   map[uint32]*Peer
	clients map[uuid.UUID]*client

	in       chan *client
	out      chan *client
	messages chan *ClientMessage
}

func NewNode(ch *Chain, capacity uint) *Node {
	node := &Node{
		chain:   ch,
		peers:   make(map[uint32]*Peer),
		clients: make(map[uuid.UUID]*client),

		in:       make(chan *client, capacity),
		out:      make(chan *client, capacity),
		messages: make(chan *ClientMessage, capacity),
	}
	return node
}

func (n *Node) AddPeer(peer *Peer) {
	n.Lock()
	defer n.Unlock()
	code := peer.HashCode()
	if _, ok := n.peers[code]; ok {
		log.Printf("peer already registered")
		return
	}
	n.peers[code] = peer
	ctx := context.TODO()
	socket, response, _ := websocket.DefaultDialer.DialContext(ctx, peer.WsUri(), http.Header{})
	fmt.Println(fmt.Sprintf("%v", response))
	configureWebsocket(socket)
	client := newClient(socket)
	log.Printf("attempting to add peer %d with address %s. Client Id: %s",
		peer.HashCode(), peer.WsUri(), client.Id)
	go n.Accept(client)
}

func (n *Node) Accept(c *client) {
	c.Join(n.in)
	defer func() { c.Leave(n.out) }()
	go c.Write()
	c.Read(n.messages)
}

func (n *Node) Run() {
	defer func() {
		n.Free()
	}()
	for {
		select {
		case client := <-n.in:
			n.clients[client.Id] = client
		case client := <-n.out:
			if _, ok := n.clients[client.Id]; ok {
				delete(n.clients, client.Id)
				client.Exit()
			}
		case msg := <-n.messages:
			n.dispatchMessage(msg)
		}
	}
}

func (n *Node) Clients() []client {
	clients := make([]client, len(n.clients))
	i := 0
	for _, client := range n.clients {
		clients[i] = *client
		i++
	}
	return clients
}

func (n *Node) BroadCast(msg Message) {
	for _, client := range n.clients {
		go client.Send(msg)
	}
}

func (n *Node) BroadCastNewBlock(block Block) {
	log.Printf("broadcasting new block %v", block)
	msg := NewBlockCreatedMessage(block)
	n.BroadCast(msg)
}

func (n *Node) Free() {
	close(n.in)
	close(n.out)
	close(n.messages)
}

func (n *Node) respond(cId uuid.UUID, msg Message) {
	client, ok := n.clients[cId]
	if !ok {
		log.Printf("there is not client to respond with id %s", cId.String())
		return
	}
	client.Send(msg)
}

func (n *Node) dispatchMessage(msg *ClientMessage) {
	switch msg.Type() {
	case BlockChainMessageRequest:
		n.handleQueryBlockchain(msg)
		return
	case BlockChainResponse:
		n.handleBlockchainReceived(msg)
		return
	case LatestBlockMessageRequest:
		n.handleQueryLatestBlock(msg)
		return
	case LatestBlockResponse:
		n.handleLatestBlockReceived(msg)
		return
	}
}
