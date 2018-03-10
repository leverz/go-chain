package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func SerializeBlock(b Block) ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		fmt.Errorf("encoding error: %s", err)
	}
	return result.Bytes(), err
}

func DeserializeBlock(d []byte) (Block, error) {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		fmt.Errorf("decoding error: %s", err)
	}
	return block, err
}
