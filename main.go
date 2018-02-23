package main

import (
	"fmt"
)

func main()  {
	blockChain := CreateBlockChain()
	blockChain.AddBlock("Send 1 BTC to Lever")
	blockChain.AddBlock("Send 0.01 BTC to Wang")

	for _, block := range blockChain.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
