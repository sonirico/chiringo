package main

import (
	"crypto/sha256"
	"encoding/hex"
)

type Hasher interface {
	Hash(Dumper) string
}

type hasherImpl struct{}

func NewHasher() Hasher {
	return &hasherImpl{}
}

func (hi *hasherImpl) Hash(dumper Dumper) string {
	h := sha256.New()
	dumper.Dump(h)
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

var hasher = NewHasher()

func Hash(dumper Dumper) string {
	return hasher.Hash(dumper)
}
