package main

import (
	"github.com/boltdb/bolt"
	"github.com/Go-zh/tools/go/gcimporter15/testdata"
	"net/mail"
)

const blocksBucket = "blocks"
const genesisCoinbaseData = "Genesis Block"

type Chain struct {
	tip []byte
	db *bolt.DB // open once use every where
}

type ChainIterator struct {
	currentHash []byte
	db *bolt.DB
} 

// 这里使用指针是因为希望直接修改 Chain 的 blocks 的值
func (c *Chain) AddBlock (data string)  {
	err := c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash := b.Get([]byte("last"))
		newBlock := NewBlock(data, lastHash)
		serializeResult, err := SerializeBlock(newBlock)
		LogError("Serialize Error", err)
		err = b.Put(newBlock.Hash, serializeResult)
		LogError("bolt put new block error", err)
		err = b.Put([]byte("last"), newBlock.Hash)
		LogError("bolt put last hash error", err)
		return nil
	})
	LogError("bolt update error", err)
}

func CreateGenesisBlock(coinbase *Transaction) Block  {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

func CreateBlockChain(address string) Chain {
	var tip []byte
	db, err := bolt.Open("chain.db", 0600, nil)
	LogError("dolt open db error", err)
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		// 数据库中没有 block
		if b == nil {
			cbtx := CreateCoinbaseTX(address, genesisCoinbaseData)
			genesis := CreateGenesisBlock(cbtx)
			b, err := tx.CreateBucket([]byte(blocksBucket))
			LogError("dolt create bucket error", err)
			serializeResult, err := SerializeBlock(genesis)
			LogError("serialize error", err)
			err = b.Put(genesis.Hash, serializeResult)
			LogError("dolt put genesis block error", err)
			err = b.Put([]byte("last"), genesis.Hash)
			LogError("bolt put last hash error", err)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("last"))
		}
		// boltDB 中 return nil 表示一个事物完成，此时会触发提交事物的操作
		return nil
	})
	LogError("blotDB createBlockChain error", err)
	return Chain{tip, db}
}

func (c *Chain) Iterator() ChainIterator {
	return ChainIterator{c.tip, c.db}
}

func (i *ChainIterator) Next() Block {
	var block Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		deserializeResult, err := DeserializeBlock(encodedBlock)
		LogError("DeserializeBlock error", err)
		block = deserializeResult

		return nil
	})

	LogError("boltDB get block error", err)

	i.currentHash = block.PrevBlockHash
	return block
}
