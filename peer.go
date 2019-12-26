package main

import (
	"fmt"
	"github.com/google/uuid"
)

type Peer struct {
	Id   uuid.UUID `json:"id"`
	Host string    `json:"host"`
	Port string    `json:"port"`
}

func NewPeer(h, p string) Peer {
	return Peer{
		Id:   uuid.New(),
		Host: h,
		Port: p,
	}
}

func (p Peer) Address() string {
	return fmt.Sprintf("%s:%s", p.Host, p.Port)
}

func (p Peer) HttpUri() string {
	return fmt.Sprintf("%s://%s", "http", p.Address())
}

func (p Peer) WsUri() string {
	return fmt.Sprintf("%s://%s", "ws", p.Address())
}
