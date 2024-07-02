package blockchain

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func Validate(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (block *Block) Serializate() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	Validate(encoder.Encode(block))

	return result.Bytes()
}

func DeserializateBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	Validate(decoder.Decode(&block))

	return &block
}
