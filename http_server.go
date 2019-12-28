package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	router.HandleFunc("/blocks", s.HandleGetBlockChain).Methods(http.MethodGet)
	router.HandleFunc("/blocks", s.HandleMineBlock).Methods(http.MethodPost)
	router.HandleFunc("/blocks/{index:[0-9]+}", s.HandleGetBlock).Methods(http.MethodGet)
	router.HandleFunc("/peers", s.HandleGetPeers).Methods(http.MethodGet)
	router.HandleFunc("/peers", s.HandleAddPeer).Methods(http.MethodPost)
	s.Handle("/", router)
}

func (s *server) Serve(port string) {
	address := fmt.Sprintf(":%s", port)
	log.Printf("HTTP server listening on %s", address)
	err := http.ListenAndServe(address, s)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
