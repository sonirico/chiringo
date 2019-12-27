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
	go s.node.BroadCastNewBlock(block)
	w.Header().Set("Location", fmt.Sprintf("/blocks/%d", block.Index))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(block)
}

type peerJson struct {
	Id   uint32 `json:"id"`
	Http string `json:"http_uri"`
	Ws   string `json:"ws_uri"`
}

func (s *server) HandleGetPeers(w http.ResponseWriter, req *http.Request) {
	peers := make([]peerJson, len(s.node.peers))
	i := 0
	for _, peer := range s.node.peers {
		peers[i] = peerJson{
			Id:   peer.HashCode(),
			Http: peer.HttpUri(),
			Ws:   peer.WsUri(),
		}
		i++
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(peers)
}

func (s *server) HandleAddPeer(w http.ResponseWriter, req *http.Request) {
	type payload struct {
		Host     string `json:"host"`
		HttpPort string `json:"http_port"`
		WsPort   string `json:"ws_port"`
	}
	var pay payload
	err := json.NewDecoder(req.Body).Decode(&pay)
	if err != nil {
		http.Error(w, "invalid peer payload", http.StatusUnprocessableEntity)
		return
	}
	pay.Host = strings.TrimSpace(pay.Host)
	pay.HttpPort = strings.TrimSpace(pay.HttpPort)
	pay.WsPort = strings.TrimSpace(pay.WsPort)
	peer := NewPeer(pay.Host, pay.HttpPort, pay.WsPort)
	if _, ok := s.node.peers[peer.HashCode()]; ok {
		http.Error(w, "peer already registered", http.StatusConflict)
		return
	}
	s.node.AddPeer(peer)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", fmt.Sprintf("/peers/%d", peer.HashCode()))
	w.WriteHeader(http.StatusCreated) // TODO: Should check connection, return 202
	jsonPeer := peerJson{
		Id:   peer.HashCode(),
		Http: peer.HttpUri(),
		Ws:   peer.WsUri(),
	}
	json.NewEncoder(w).Encode(jsonPeer)
}
