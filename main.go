package main

import (
	"fmt"

	"github.com/aureuneun/bitcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	chain.AddBlock("Second BLock")
	chain.AddBlock("Third Block")
	for _, block := range chain.AllBlocks() {
		fmt.Println(block)
	}
}
