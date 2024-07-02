package blockchain

import "time"

// Block
type Block struct {
	Timestamp int64
	Hash      []byte
	Data      []byte
	PrevHash  []byte
	Nonce     int64
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte{}, []byte(data), prevHash, 0}
	proofOfWork := NewProofOfWork(block)
	nonce, hash := proofOfWork.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

func GenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte(""))
}
