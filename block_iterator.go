package main

import (
	"bytes"
	"encoding/gob"

	"github.com/boltdb/bolt"
)

type BlockIterator struct {
	CurrentHash []byte
	DB          *DB
	Data        []byte
}

func NewBlockIterator(db *DB, currentHash []byte) *BlockIterator {
	return &BlockIterator{
		CurrentHash: currentHash,
		DB:          db,
	}
}

func (bi *BlockIterator) Decode(dest *Block) {
	r := bytes.NewReader(bi.Data)
	err := gob.NewDecoder(r).Decode(dest)
	if err != nil {
		panic(err)
	}
	bi.CurrentHash = dest.PreBlockHash
}

func (bi *BlockIterator) Next() bool {
	var data []byte
	bi.DB.Bolt.View(func(tx *bolt.Tx) error {
		bucket := bi.DB.Bucket(tx)
		data = bucket.Get(bi.CurrentHash)
		return nil
	})
	if data == nil {
		return false
	}

	bi.Data = data
	return true
}
