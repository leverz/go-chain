package main

//import (
//	"fmt"
//	"strconv"
//)

func main()  {
	blockChain := CreateBlockChain()
	defer blockChain.db.Close()

	cli := CLI{&blockChain}
	cli.Run()

	//for _, block := range blockChain.blocks {
	//	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	//	fmt.Printf("Data: %s\n", block.Data)
	//	fmt.Printf("Hash: %x\n", block.Hash)
	//	pow := createProofOfWork(block)
	//	fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
	//	fmt.Println()
	//}
}
