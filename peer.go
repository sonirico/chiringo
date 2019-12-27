package main

import (
	"fmt"
	"hash/fnv"
)

type Peer struct {
	Host     string `json:"host"`
	HttpPort string `json:"http_port"`
	WsPort   string `json:"ws_port"`
}

func NewPeer(host, httpPort, wsPort string) *Peer {
	return &Peer{
		Host:     host,
		HttpPort: httpPort,
		WsPort:   wsPort,
	}
}

func (p Peer) HttpUri() string {
	return fmt.Sprintf("%s://%s:%s", "http", p.Host, p.HttpPort)
}

func (p Peer) WsUri() string {
	return fmt.Sprintf("%s://%s:%s", "ws", p.Host, p.WsPort)
}

func (p Peer) HashCode() uint32 {
	payload := fmt.Sprintf("%s%s%s", p.Host, p.HttpPort, p.WsPort)
	h := fnv.New32a()
	h.Write([]byte(payload))
	return h.Sum32()
}

type PeerClient struct {
	peer *Peer
}
