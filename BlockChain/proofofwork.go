package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

// Proof Of Work
const targetBits = 20
const maxNonce = math.MaxInt64

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	return &ProofOfWork{block: block, target: target}
}

func (proofOfWork *ProofOfWork) prepareData(nounce int64) []byte {
	data := bytes.Join(
		[][]byte{
			proofOfWork.block.PrevHash,
			proofOfWork.block.Data,
			IntToHex(proofOfWork.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nounce)),
		}, []byte{},
	)

	return data
}

func (proofOfWork *ProofOfWork) Validate() bool {
	var hasInt big.Int

	data := proofOfWork.prepareData(proofOfWork.block.Nonce)
	hash := sha256.Sum256(data)
	hasInt.SetBytes(hash[:])

	isValid := hasInt.Cmp(proofOfWork.target) == -1
	return isValid
}

func (proofOfWork *ProofOfWork) Run() (int64, []byte) {
	var hashInt big.Int
	var hash [32]byte
	var nonce int64 = 0

	fmt.Printf("Mining the block containing \"%s\"\n", proofOfWork.block.Data)
	for nonce < maxNonce {
		data := proofOfWork.prepareData(nonce)
		hash = sha256.Sum256(data)
		//fmt.Printf("\r%x", hash)

		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(proofOfWork.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Printf("\n\n")

	return nonce, hash[:]
}
