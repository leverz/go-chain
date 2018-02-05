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
	Timestamp int64 // Block 创建时的时间戳
	Data []byte // Block 中存储的有意义的信息
	PrevBlockHash []byte // 上一个区块的 Hash 值
	Hash []byte  // 本区块的 Hash 值
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10)) // int64 转为 字符串
	headers := bytes.Join([][]byte{ b.PrevBlockHash, b.Data, timestamp }, []byte{})
	hash := sha256.Sum256(headers)
  // 这里似乎传递的是 hash 的引用
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}
