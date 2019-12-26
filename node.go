package main

import "sync"

type Node struct {
	sync.Mutex

	chain *Chain
	peers []Peer
}

func NewNode(ch *Chain) *Node {
	return &Node{
		chain: ch,
	}
}

func (n *Node) AddPeer(peer Peer) {
	n.Lock()
	n.peers = append(n.peers, peer)
	n.Unlock()
}
