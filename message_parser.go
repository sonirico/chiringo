package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

type messagePayload struct {
	Type MessageType `json:"code"`
}

type MessageParserFunc func(reader io.Reader) (Message, error)

type MessageParser struct {
	reg map[MessageType]MessageParserFunc
}

func NewMessageParser() *MessageParser {
	parser := &MessageParser{
		reg: make(map[MessageType]MessageParserFunc),
	}
	parser.register(BlockChainMessageRequest, parser.parseGetBlockchainRequest)
	parser.register(LatestBlockMessageRequest, parser.parseGetLatestBlockRequest)
	parser.register(BlockChainResponse, parser.parseGetBlockchainResponse)
	parser.register(LatestBlockResponse, parser.parseGetLatestBlockResponse)
	return parser
}

func (m *MessageParser) register(mType MessageType, parserFunc MessageParserFunc) {
	m.reg[mType] = parserFunc
}

func (m *MessageParser) parseGetBlockchainResponse(reader io.Reader) (Message, error) {
	response := &MessageResponseBlockchain{}
	err := json.NewDecoder(reader).Decode(response)
	return response, err
}

func (m *MessageParser) parseGetLatestBlockResponse(reader io.Reader) (Message, error) {
	var p struct {
		Message struct {
			Block Block `json:"block"`
		} `json:"message"`
	}
	err := json.NewDecoder(reader).Decode(&p)
	if err != nil {
		return nil, err
	}
	msg := NewBlockCreatedMessage(p.Message.Block)
	return msg, nil
}

func (m *MessageParser) parseGetBlockchainRequest(reader io.Reader) (Message, error) {
	request := &MessageRequestBlockchain{}
	err := json.NewDecoder(reader).Decode(request)
	return request, err
}

func (m *MessageParser) parseGetLatestBlockRequest(reader io.Reader) (Message, error) {
	request := &MessageRequestLatestBlock{}
	err := json.NewDecoder(reader).Decode(request)
	return request, err
}

func (m *MessageParser) Parse(data []byte) (Message, error) {
	var result messagePayload
	buf := bytes.NewReader(data)
	err := json.NewDecoder(buf).Decode(&result)
	if err != nil {
		return nil, err
	}
	if _, err = buf.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}
	fn, ok := m.reg[result.Type]
	if !ok {
		return nil, errors.New("unknown message type")
	}
	return fn(buf)
}
