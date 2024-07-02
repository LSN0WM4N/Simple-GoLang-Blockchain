package blockchain

import (
	"github.com/boltdb/bolt"
)

const dbFile = "database.db"

type BlockChain struct {
	tip []byte
	Db  *bolt.DB
}

type BlockChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (blockChain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	Validate(blockChain.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		lastHash = b.Get([]byte("l"))

		return nil
	}))

	newBlock := NewBlock(data, lastHash)

	Validate(blockChain.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		b.Put(newBlock.Hash, newBlock.Serializate())
		b.Put([]byte("l"), newBlock.Hash)

		blockChain.tip = newBlock.Hash

		return nil
	}))
}

func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	Validate(err)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))

		if b == nil { // Empty BlockChain
			genesis := GenesisBlock()
			b, _ := tx.CreateBucket([]byte("blocks"))
			b.Put(genesis.Hash, genesis.Serializate())
			b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	blockChain := BlockChain{tip, db}

	return &blockChain
}

func (blockChain *BlockChain) Iterator() *BlockChainIterator {
	blockChainIterator := &BlockChainIterator{blockChain.tip, blockChain.Db}
	return blockChainIterator
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block

	Validate(iter.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blocks"))
		encodedBlock := b.Get(iter.currentHash)
		block = DeserializateBlock(encodedBlock)

		return nil
	}))

	iter.currentHash = block.PrevHash

	return block
}
