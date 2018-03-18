package main

import (
	"fmt"
	"bytes"
	"encoding/gob"
	"crypto/sha256"
)

// 挖矿奖励的数量, 目前暂时设计成固定的，之后可以模仿比特币动态计算
const subsidy = 10

type Transaction struct {
	ID []byte
	Input []TXInput
	Output []TXOutput
}

type TXInput struct {
	Txid []byte // 所属的交易 id
	Output int // 存储序号，一个交易有可能有多个 TXO
	ScriptSig string // 钱包的密钥
}

type TXOutput struct {
	Value int // 交易的数量
	ScriptPubKey string // 钱包公钥
}

func (tx *Transaction) SetID()  {
	var encoded bytes.Buffer
	var hash [32]byte

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	LogError("transaction SetID error", err)
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func CreateCoinbaseTX(to, data string) *Transaction  {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()

	return &tx
}
