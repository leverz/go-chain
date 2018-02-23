package main

type Chain struct {
	blocks []Block
}

// 这里使用指针是因为希望直接修改 Chain 的 blocks 的值
func (c *Chain) AddBlock (data string)  {
	prevBlock := c.blocks[len(c.blocks) - 1]
	newBlock := NewBlock(data, prevBlock.Hash)
	c.blocks = append(c.blocks, newBlock)
}

func CreateGenesisBlock() Block  {
	return NewBlock("Genesis Block", []byte{})
}

func CreateBlockChain() Chain {
	return Chain{[]Block{CreateGenesisBlock()}}
}
