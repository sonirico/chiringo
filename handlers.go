package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (n *Node) GetBlockChain(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(n.chain.elements)
}

func (n *Node) GetBlock(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	blockIndexParam := vars["index"]
	index, err := strconv.Atoi(blockIndexParam)
	if err != nil || index < 0 || index >= n.chain.Size() {
		http.NotFound(w, req)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(n.chain.elements[index])
}

func (n *Node) MineBlock(w http.ResponseWriter, req *http.Request) {
	type payload struct {
		Data string `json:"data"`
	}
	var pay payload
	err := json.NewDecoder(req.Body).Decode(&pay)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusUnprocessableEntity)
		return
	}
	block := n.chain.NextBlock([]byte(pay.Data))
	w.Header().Set("Location", fmt.Sprintf("%s/blocks/%d", req.Host, block.Index))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(block)
}
