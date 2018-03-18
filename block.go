package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

/**
 * 在比特币中，Timestamp、Data、PrevBlockHash 组成一个数据结构(块头)，而 Data 对应一个单独的数据结构
 * go 语言中希望暴露给外部的方法首字母要大写
 **/

// 区块的数据结构
type Block struct {
	Timestamp int64  // Block 创建时的时间戳
	Transactions []*Transaction // Block 中存储的有意义的信息
	PrevBlockHash []byte // 上一个区块的 Hash 值
	Hash          []byte // 本区块的 Hash 值
	Nonce int
}

// 这里使用指针类型是希望调用该方法直接修改 Block 中的 Hash 字段的值
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10)) // int64 转为 字符串
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func (b *Block) HashTransactions() []byte  {
	var txHashes [][]byte
	var txHash [32]byte
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

func NewBlock(txs []*Transaction, prevBlockHash []byte) Block {
	block := Block{time.Now().Unix(), txs, prevBlockHash, []byte{}, 0}
	pow := createProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}
