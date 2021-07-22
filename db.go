package main

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var (
	lastKey    = []byte("l")
	bucketName = []byte("bucket_1")
)

const (
	dbFile = "blockchain.db"
)

type DB struct {
	Bolt *bolt.DB
}

func NewDB() (*DB, func()) {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		panic(err)
	}

	return &DB{Bolt: db}, func() {
		err = db.Close()
		if err != nil {
			fmt.Printf("close db failed: %v\n", err)
		}
	}
}

func (d *DB) Bucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket(bucketName)
}
