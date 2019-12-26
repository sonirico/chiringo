package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func (s *server) HandleGetBlockChain(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s.node.chain.elements)
}

func (s *server) HandleGetBlock(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	blockIndexParam := vars["index"]
	index, err := strconv.Atoi(blockIndexParam)
	if err != nil || index < 0 || index >= s.node.chain.Size() {
		http.NotFound(w, req)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s.node.chain.elements[index])
}

func (s *server) HandleMineBlock(w http.ResponseWriter, req *http.Request) {
	type payload struct {
		Data string `json:"data"`
	}
	var pay payload
	err := json.NewDecoder(req.Body).Decode(&pay)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusUnprocessableEntity)
		return
	}
	block := s.node.chain.NextBlock([]byte(pay.Data))
	w.Header().Set("Location", fmt.Sprintf("/blocks/%d", block.Index))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(block)
}

type peerJson struct {
	Peer
	Http string `json:"http_uri"`
	Ws   string `json:"ws_uri"`
}

func (s *server) GetPeers(w http.ResponseWriter, req *http.Request) {
	peers := make([]peerJson, len(s.node.peers))
	for i, peer := range s.node.peers {
		peers[i] = peerJson{
			Peer: peer,
			Http: peer.HttpUri(),
			Ws:   peer.WsUri(),
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peers)
}

func (s *server) AddPeer(w http.ResponseWriter, req *http.Request) {
	type payload struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	var pay payload
	err := json.NewDecoder(req.Body).Decode(&pay)
	if err != nil {
		http.Error(w, "invalid peer payload", http.StatusUnprocessableEntity)
		return
	}
	pay.Host = strings.TrimSpace(pay.Host)
	pay.Port = strings.TrimSpace(pay.Port)
	// TODO: Perform further checks on host-port pairs
	for _, peer := range s.node.peers {
		if peer.Host == pay.Host && pay.Port == peer.Port {
			http.Error(w, "peer already registered", http.StatusConflict)
			return
		}
	}
	peer := NewPeer(pay.Host, pay.Port)
	s.node.AddPeer(peer)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", fmt.Sprintf("/peers/%s", peer.Id))
	w.WriteHeader(http.StatusCreated)
	jsonPeer := peerJson{
		Peer: peer,
		Http: peer.HttpUri(),
		Ws:   peer.WsUri(),
	}
	json.NewEncoder(w).Encode(jsonPeer)
}
