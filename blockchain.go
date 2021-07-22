package main

import (
	"fmt"

	"github.com/boltdb/bolt"
)

type BlockChain struct {
	Tip []byte
	DB  *DB
}

func NewBlockChain() (*BlockChain, func()) {
	db, closeDB := NewDB()

	var (
		tip []byte
		err = db.Bolt.Update(func(tx *bolt.Tx) error {
			bucket := db.Bucket(tx)

			if bucket != nil {
				tip = bucket.Get(lastKey)
				if tip == nil {
					panic("invalid tip")
				}

				return nil
			}

			fmt.Printf("not existsing blockchain found. Creating...\n")

			bucket, err := tx.CreateBucket(bucketName)
			if err != nil {
				panic(err)
			}

			orgNode := NewOriginNode()
			err = bucket.Put(orgNode.Hash, orgNode.Serialize())
			if err != nil {
				panic(err)
			}

			err = bucket.Put(lastKey, orgNode.Hash)
			if err != nil {
				panic(err)
			}

			tip = orgNode.Hash

			return nil
		})
	)

	if err != nil {
		panic(err)
	}

	fmt.Printf("access into blockchain successfully\n\n")
	return &BlockChain{
		Tip: tip,
		DB:  db,
	}, closeDB
}

func (bc *BlockChain) AddBlock(data []byte) {
	var lastHash []byte
	err := bc.DB.Bolt.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		lastHash = bucket.Get(lastKey)
		return nil
	})
	if err != nil {
		panic(err)
	}

	bl := NewBlock(data, lastHash)
	err = bc.DB.Bolt.Update(func(tx *bolt.Tx) error {
		bucket := bc.DB.Bucket(tx)
		bl.AddToBucket(bucket)

		err = bucket.Put(lastKey, bl.Hash)
		if err != nil {
			panic(err)
		}

		bc.Tip = bl.Hash
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func (bc *BlockChain) Print() {
	iter := NewBlockIterator(bc.DB, bc.Tip)
	for iter.Next() {
		dest := &Block{}
		iter.Decode(dest)
		dest.PrintBeauty()
	}
}
