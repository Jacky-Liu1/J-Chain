package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Method for generating a hash for the current block
func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10)) // array of digits
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.AllData}, []byte{})
	hash := sha256.Sum256(headers)
	block.MyBlockHash = hash[:]
}

// Generate and return a new block
func NewBlock(data string, PreviousBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), PreviousBlockHash, []byte{}, []byte(data)}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis block", []byte{})
}
