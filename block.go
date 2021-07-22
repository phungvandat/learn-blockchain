package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

type Block struct {
	Timestamp    int64
	Data         []byte
	PreBlockHash []byte
	Hash         []byte
	Nonce        int
}

func NewBlock(data, preBlockHash []byte) *Block {
	bl := &Block{
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PreBlockHash: preBlockHash,
	}

	pow := NewProofOfWork(bl)
	bl.Nonce, bl.Hash = pow.Run()

	return bl
}

// SetHash example samplest to create hash for new block
func (bl *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(bl.Timestamp, 10))
	bodyHash := bytes.Join([][]byte{bl.Data, timestamp, bl.PreBlockHash}, nil)
	currentHash := sha256.Sum256(bodyHash)
	bl.Hash = currentHash[:]
}

func NewOriginNode() *Block {
	return NewBlock([]byte("init block"), nil)
}

// Serialize encode block to []byte
func (bl *Block) Serialize() []byte {
	buf := &bytes.Buffer{}

	err := gob.NewEncoder(buf).Encode(bl)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func (bl *Block) AddToBucket(bk *bolt.Bucket) {
	data := bl.Serialize()
	err := bk.Put(bl.Hash, data)
	if err != nil {
		panic(err)
	}
}

func (bl Block) PrintBeauty() {
	m := map[string]interface{}{
		"time":           time.Unix(bl.Timestamp, 0),
		"data":           string(bl.Data),
		"hash":           fmt.Sprintf("%x", bl.Data),
		"pre_block_hash": fmt.Sprintf("%x", bl.PreBlockHash),
		"nonce":          bl.Nonce,
	}
	PrintJSON(m)
}
