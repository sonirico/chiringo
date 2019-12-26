package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"io"
	"time"
)

type Dumper interface {
	Dump(writer io.Writer)
}

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
	ts := time.Date(2019, 12, 25, 16, 10, 00, 0, time.UTC)
	block := NewBlock(0, ts, "", "", []byte("Sitting in the dock of the bay"))
	block.Hash = Hash(block)
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

func (b Block) Dump(w io.Writer) {
	indexBuf := make([]byte, 8)
	binary.LittleEndian.PutUint64(indexBuf, b.Index)
	_, _ = w.Write(indexBuf)
	_, _ = w.Write([]byte(b.PreviousHash))
	_, _ = w.Write([]byte(b.Timestmap.String()))
	_, _ = w.Write(b.Data)
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
