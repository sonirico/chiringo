package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"io"
	"time"
)

type Block struct {
	Index        uint64    `json:"index"`
	Hash         string    `json:"hash"`
	PreviousHash string    `json:"prev_hash"`
	Data         []byte    `json:"data"`
	Timestmap    time.Time `json:"time"`
}

func NewBlock(index uint64, timestamp time.Time, hash, previousHash string, data []byte) Block {
	return Block{
		Index:        index,
		Hash:         hash,
		PreviousHash: previousHash,
		Data:         data,
		Timestmap:    timestamp,
	}
}

func NewGenesisBlock() Block {
	ts := time.Date(2019, time.December, 25, 16, 10, 00, 0, time.UTC)
	block := NewBlock(0, ts, "", "", []byte("Sitting in the dock of the bay"))
	block.Hash = block.CalculateHash()
	return block
}

func (b Block) String() string {
	var buf bytes.Buffer
	b.Dump(&buf)
	return buf.String()
}

func (b Block) Bytes() []byte {
	var buf bytes.Buffer
	b.Dump(&buf)
	return buf.Bytes()
}

func (b Block) Dump(buf io.Writer) {
	indexBuf := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexBuf, b.Index)
	buf.Write(indexBuf)
	buf.Write([]byte(b.PreviousHash))
	buf.Write([]byte(b.Timestmap.String()))
	buf.Write(b.Data)
}

func (b Block) CalculateHash() string {
	h := sha256.New()
	indexBuf := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexBuf, b.Index)
	b.Dump(h)
	return string(h.Sum(nil))
}

type JSONBlock struct {
	Block
}

func (jb *JSONBlock) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, jb)
}

func (jb *JSONBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(jb.Block)
}
