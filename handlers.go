package main

import (
	"encoding/json"
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
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(n.chain.elements[index])
}
