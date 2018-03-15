package main

import (
	"os"
	"fmt"
	"strconv"
	"flag"
)

type CLI struct {
	c *Chain
}

func (cli *CLI) printUsage()  {
	fmt.Printf("Usage: ")
	fmt.Printf("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Printf("  printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs()  {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string)  {
	cli.c.AddBlock(data)
	fmt.Printf("Success!")
}

func printChain(i ChainIterator) {
	block := i.Next()
	fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Hash: %x\n", block.Hash)
	pow := createProofOfWork(block)
	fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
	fmt.Println()
	if len(block.PrevBlockHash) > 0 {
		printChain(i)
	}
}

func (cli *CLI) printChain()  {
	i := cli.c.Iterator()
	printChain(i)
}

func (cli *CLI) Run()  {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		LogError("CMD parse error", err)
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		LogError("CMD parse error", err)
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		// 输入参数为空，显示提示信息并推出
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
