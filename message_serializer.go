package main

import (
	"bytes"
	"encoding/json"
)

type serializeResult struct {
	Message Message     `json:"message"`
	Type    MessageType `json:"code"`
}

func Serialize(msg Message) []byte {
	var buf bytes.Buffer
	res := &serializeResult{
		Type:    msg.Type(),
		Message: msg,
	}
	err := json.NewEncoder(&buf).Encode(res)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}
