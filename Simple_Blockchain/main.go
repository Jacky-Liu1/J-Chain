package main

import (
	"fmt"
)

func main() {
	newBlockchain := NewBlockchain()
	newBlockchain.AddBlock("first transaction")
	newBlockchain.AddBlock("second transaction")
	for _, block := range newBlockchain.Blocks {
		fmt.Printf("Hash of the block %x\n", block.MyBlockHash)                 // print the hash of the block
		fmt.Printf("Hash of the previous Block: %x\n", block.PreviousBlockHash) // print the hash of the previous block
		fmt.Printf("All the transactions: %s\n", block.AllData)
	}
}
