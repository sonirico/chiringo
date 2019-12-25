package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Peer struct {
	host string
	port string
}

type Node struct {
	chain     *Chain
	peers     []Peer
	webServer *http.ServeMux
}

func NewNode(ch *Chain) *Node {
	node := &Node{
		chain: ch,
	}
	webServer := http.NewServeMux()
	webServer.HandleFunc("/", node.GetBlockChain)
	node.webServer = webServer
	return node
}

func (n *Node) ServeHTTP() {
	port := os.Getenv("HTTP_PORT")
	address := fmt.Sprintf(":%s", port)
	log.Printf("HTTP server listening on %s", address)
	err := http.ListenAndServe(address, n.webServer)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
