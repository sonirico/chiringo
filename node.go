package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type Peer struct {
	Host string `json:"host"`
	Port string `json:"port"`
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

type Node struct {
	chain     *Chain
	peers     []Peer
	webServer *http.ServeMux
}

func NewNode(ch *Chain) *Node {
	node := &Node{
		chain: ch,
	}
	// Web server
	webServer := http.NewServeMux()
	router := mux.NewRouter()
	router.HandleFunc("/blocks", node.GetBlockChain).Methods(http.MethodGet)
	router.HandleFunc("/blocks", node.MineBlock).Methods(http.MethodPost)
	router.HandleFunc("/blocks/{index:[0-9]+}", node.GetBlock).Methods(http.MethodGet)
	router.HandleFunc("/peers", node.GetPeers).Methods(http.MethodGet)
	webServer.Handle("/", router)
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
