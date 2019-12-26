package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type server struct {
	http.ServeMux

	node *Node
}

func newServer(node *Node) *server {
	return &server{
		node: node,
	}
}

func (s *server) setUp() {
	router := mux.NewRouter()
	router.HandleFunc("/blocks", s.GetBlockChain).Methods(http.MethodGet)
	router.HandleFunc("/blocks", s.MineBlock).Methods(http.MethodPost)
	router.HandleFunc("/blocks/{index:[0-9]+}", s.GetBlock).Methods(http.MethodGet)
	router.HandleFunc("/peers", s.GetPeers).Methods(http.MethodGet)
	router.HandleFunc("/peers", s.AddPeer).Methods(http.MethodPost)
	s.Handle("/", router)
}

func (s *server) Serve() {
	port := os.Getenv("HTTP_PORT")
	address := fmt.Sprintf(":%s", port)
	log.Printf("HTTP server listening on %s", address)
	err := http.ListenAndServe(address, s)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
