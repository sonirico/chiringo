package main

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
)

type MessageType uint

const (
	LatestBlockMessageRequest MessageType = iota + 1
	LatestBlockResponse
	BlockChainMessageRequest
	BlockChainResponse
)

type Message interface {
	Type() MessageType
	Bytes() []byte
	String() string
}

type MessageRequestBlockchain struct{}

func (m *MessageRequestBlockchain) Type() MessageType {
	return BlockChainMessageRequest
}

func (m *MessageRequestBlockchain) Bytes() []byte {
	return []byte("GET_ALL_BLOCKS")
}

func (m *MessageRequestBlockchain) String() string {
	return string(m.Bytes())
}

type MessageRequestLatestBlock struct{}

func (m *MessageRequestLatestBlock) Type() MessageType {
	return LatestBlockMessageRequest
}
func (m *MessageRequestLatestBlock) Bytes() []byte {
	return []byte("LATEST_BLOCK")
}
func (m *MessageRequestLatestBlock) String() string {
	return string(m.Bytes())
}

type MessageResponseLatestBlock struct {
	Block Block `json:"block"`
}

func NewBlockCreatedMessage(block Block) *MessageResponseLatestBlock {
	return &MessageResponseLatestBlock{Block: block}
}
func (m *MessageResponseLatestBlock) Type() MessageType {
	return LatestBlockResponse
}
func (m *MessageResponseLatestBlock) Bytes() []byte {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(m); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
func (m *MessageResponseLatestBlock) String() string {
	return string(m.Bytes())
}

type MessageResponseBlockchain struct {
	Blocks []Block `json:"blocks"`
}

func NewBlockchainMessage(blocks []Block) *MessageResponseBlockchain {
	return &MessageResponseBlockchain{
		Blocks: blocks,
	}
}

func (m *MessageResponseBlockchain) Type() MessageType {
	return BlockChainResponse
}

func (m *MessageResponseBlockchain) Bytes() []byte {

	return []byte("BLOCKCHAIN")
}
func (m *MessageResponseBlockchain) String() string {
	return string(m.Bytes())
}

// ClientMessage holds any message along with the client Id
// which requested information
type ClientMessage struct {
	msg      Message
	clientId uuid.UUID
}

func NewClientMessage(id uuid.UUID, msg Message) *ClientMessage {
	return &ClientMessage{
		msg:      msg,
		clientId: id,
	}
}
func (m *ClientMessage) Type() MessageType {
	return m.msg.Type()
}
func (m *ClientMessage) Bytes() []byte {
	var buf bytes.Buffer
	buf.Write(m.msg.Bytes())
	buf.Write([]byte(m.clientId.String()))
	return buf.Bytes()
}
func (m *ClientMessage) String() string {
	return string(m.Bytes())
}

func (m *ClientMessage) ClientId() uuid.UUID {
	return m.clientId
}
