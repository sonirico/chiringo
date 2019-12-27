package main

import "log"

func (n *Node) handleBlockchainReceived(msg *ClientMessage) {
	m, _ := msg.msg.(*MessageResponseBlockchain)
	receivedChainSize := len(m.Blocks)
	if receivedChainSize < 1 {
		log.Printf("recevied empty chain...?")
		return
	}
	latestBlockReceived := m.Blocks[receivedChainSize-1]
	latestCurrent := n.chain.last
	if latestBlockReceived.Index > latestCurrent.Index {
		log.Printf("chain possibly behind. Current index is %d, got %d", latestCurrent.Index, latestBlockReceived.Index)
		if receivedChainSize == 1 {
			n.BroadCast(&MessageRequestBlockchain{})
		} else if latestBlockReceived.PreviousHash == latestCurrent.Hash {
			if blockIsValid(latestBlockReceived, *latestCurrent) {
				n.chain.Append(latestBlockReceived)
				// TODO: Use context to cancel requests after the first response has
				// arrived
				go n.BroadCast(&MessageRequestLatestBlock{})
			}
		} else {
			// Received blockchain is larger than current
			n.chain.Replace(m.Blocks)
		}
	} else {
		log.Printf("recevied chain is smaller than current")
		go n.respond(msg.ClientId(), &MessageResponseBlockchain{Blocks: n.chain.elements})
	}
}

func (n *Node) handleLatestBlockReceived(msg *ClientMessage) {
	m, _ := msg.msg.(*MessageResponseLatestBlock)
	receivedIndex := m.Block.Index
	currentIndex := n.chain.last.Index
	log.Printf("candidate latest block received %v", m.Block)
	if receivedIndex == currentIndex+1 {
		// Is this block the next?
		log.Printf("chain possibly behind. Current index is %d, got %d", currentIndex, receivedIndex)
		if blockIsValid(m.Block, *n.chain.last) {
			n.chain.Append(m.Block)
			// Inform peers that habemus new block!
			n.BroadCastNewBlock(*n.chain.last)
		} else {
			log.Printf("invalid block! %v. Fork or hacking attempt!...", m.Block)
		}
	} else if receivedIndex > currentIndex {
		// We are far behind. Request whole blockchain
		// THOUGHT: Use context to cancel requests after the first response has
		// arrived?
		n.BroadCast(&MessageRequestBlockchain{})
	} else {
		log.Printf("block is behind. update peer blockchain")
		blockChainMsg := NewBlockchainMessage(n.chain.Blocks())
		n.respond(msg.clientId, blockChainMsg)
	}
}

func (n *Node) handleQueryBlockchain(msg *ClientMessage) {
	response := &MessageResponseBlockchain{Blocks: n.chain.elements}
	go n.respond(msg.ClientId(), response)
}

func (n *Node) handleQueryLatestBlock(msg *ClientMessage) {
	response := &MessageResponseLatestBlock{Block: n.chain.LatestBlock()}
	go n.respond(msg.ClientId(), response)
}
