package main

import (
	"log"
	"sync"
	"time"
)

type Chain struct {
	sync.RWMutex

	genesis  Block
	last     *Block
	elements []Block
}

func NewChain() *Chain {
	c := &Chain{
		genesis: NewGenesisBlock(),
		last:    nil,
	}
	c.last = &c.genesis
	c.Append(c.genesis)
	return c
}

func (c *Chain) Append(b Block) {
	c.RLock()
	defer c.RUnlock()
	c.elements = append(c.elements, b)
}

func (c *Chain) NextBlock(data []byte) Block {
	var prevBlock Block
	if c.last != nil {
		prevBlock = *c.last
	} else {
		prevBlock = c.genesis
	}
	now := time.Now()
	newBlock := NewBlock(prevBlock.Index+1, now, "", prevBlock.Hash, data)
	newBlock.Hash = Hash(newBlock)
	c.last = &newBlock
	c.Append(newBlock)
	return newBlock
}

func (c *Chain) IsValid() bool {
	return chainIsValid(c.elements)
}

func (c *Chain) Blocks() []Block {
	c.Lock()
	defer c.Unlock()
	blocks := make([]Block, len(c.elements))
	copy(blocks, c.elements)
	return blocks
}

func (c Chain) Size() int {
	return len(c.elements)
}

func (c *Chain) Replace(blocks []Block) {
	if !chainIsValid(blocks) {
		log.Fatal("Chain is invalid, unable to replace")
		return
	}
	if c.Size() > len(blocks) {
		log.Fatal("current is Chain is larger than received, unable to replace")
		return
	}
	c.RLock()
	defer c.RUnlock()
	c.elements = blocks
	c.genesis = blocks[0]
	c.last = &blocks[len(blocks)-1]
}

func chainIsValid(blocks []Block) bool {
	genesis := blocks[0]
	if genesis.Hash != Hash(genesis) {
		log.Fatal("genesis block compromised")
		return false
	}
	end := len(blocks)
	index := 1
	for index < end {
		prev := blocks[index-1]
		next := blocks[index]
		if !blockIsValid(next, prev) {
			log.Fatalf("Chain got invalid at index: %d", index)
			return false
		}
		index++
	}
	return true
}

func (c *Chain) LatestBlock() Block {
	return *c.last
}

func blockIsValid(next, prev Block) bool {
	if prev.Index+1 != next.Index {
		log.Fatalf("expected index %d, got %d", prev.Index+1, next.Index)
		return false
	}

	if prev.Hash != next.PreviousHash {
		log.Fatal("previous hash mismatch")
		return false
	}

	if next.Hash != Hash(next) {
		log.Fatal("current hash mismatch")
		return false
	}
	return true
}
