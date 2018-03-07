package main

import (
	"math/big"
	"fmt"
	"bytes"
	"math"
	"crypto/sha256"
	"strconv"
)

// 比特币中使用 Hashcash 算法计算 Hash

// 比特币中 target bits 是区块的 header 信息中存储的表示区块开采难度的字段, 在比特币中会动态调整开采难度的算法
const targetBits = 24

const maxNonce = math.MaxInt64

type ProofOfWork struct {
	block Block
	target big.Int
}

func IntToHex(data int64) []byte {
	return []byte( strconv.FormatInt(data, 16) )
}

func createProofOfWork(b Block) ProofOfWork  {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBits))

	return ProofOfWork{b, *target}
}

func (pow ProofOfWork) prepareData(nonce int) []byte  {
	return bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
}

func (pow ProofOfWork) Run() (int, []byte)  {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Missing the block containing \"%s\"\n", pow.block.Data)

	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(&pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")
	return nonce, hash[:]
}
