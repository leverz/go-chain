package main

import (
	"bytes"
	"encoding/gob"
)

func SerializeBlock(b Block) ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	LogError("encoding error", err)
	return result.Bytes(), err
}

func DeserializeBlock(d []byte) (Block, error) {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	LogError("decoding error", err)
	return block, err
}
