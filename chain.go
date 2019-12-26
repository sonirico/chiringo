package main

import (
	"log"
	"time"
)

type Chain struct {
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
	c.elements = append(c.elements, b)
}

func (c *Chain) NextBlock(data []byte) Block {
	var b Block
	if c.last != nil {
		b = *c.last
	} else {
		b = c.genesis
	}
	now := time.Now()
	newBlock := NewBlock(b.Index+1, now, "", b.Hash, data)
	newBlock.Hash = Hash(newBlock)
	c.last = &newBlock
	c.Append(newBlock)
	return newBlock
}

func (c *Chain) IsValid() bool {
	return chainIsValid(c.elements)
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
