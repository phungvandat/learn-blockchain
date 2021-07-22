package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

const targetBits uint = 24

func NewProofOfWork(bl *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{
		Block:  bl,
		Target: target,
	}

	return pow
}

func (p *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		p.Block.PreBlockHash,
		p.Block.Data,
		[]byte(fmt.Sprintf("%x", targetBits)),
		[]byte(fmt.Sprintf("%x", nonce)),
	}, nil)
	return data
}

func (p *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("mining the block with data: \"%s\"\n", p.Block.Data)

	for nonce < math.MaxInt64 {
		data := p.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(p.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\nsuccess with nonce: %v\n\n", nonce)
	return nonce, hash[:]
}

func (p *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := p.prepareData(p.Block.Nonce)

	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(p.Target) == -1
}
