package main

import (
	"encoding/json"
	"net/http"
)

func (n *Node) GetBlockChain(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(n.chain.elements)
}
