package main

type Block struct {
	Timestamp     int64  // Time when the block was created
	Data          []byte // Actual valuable information
	PrevBlockHash []byte // Stores the hash of the previous block
	Hash          []byte // Hash of the block
}
