package main

import (
	"crypto/sha256"
	"encoding/binary"
)

type BlockHeader struct {
	Version           uint32   // 0..3
	PreviousBlockHash [32]byte // 4..35
	MerkleRoot        [32]byte // 36..67
	Timestamp         uint32   // 68..71
	Difficulty        uint32   // 72..75
	Nonce             uint32   // 76..79
}

func (bh *BlockHeader) Encode() (header [80]byte) {
	binary.BigEndian.PutUint32(header[0:], bh.Version)
	copy(header[4:], bh.PreviousBlockHash[:])
	copy(header[36:], bh.MerkleRoot[:])
	binary.BigEndian.PutUint32(header[68:], bh.Timestamp)
	binary.BigEndian.PutUint32(header[72:], bh.Difficulty)
	binary.BigEndian.PutUint32(header[76:], bh.Nonce)
	return
}

type Transaction struct {
	From   [20]byte // 0..19
	To     [20]byte // 20..39
	Amount uint32   // 40..43
}

func (tx *Transaction) Encode() (transaction [44]byte) {
	copy(transaction[0:], tx.From[:])
	copy(transaction[20:], tx.To[:])
	binary.BigEndian.PutUint32(transaction[40:], tx.Amount)
	return
}

func (tx *Transaction) ID() [32]byte {
	etx := tx.Encode()
	return sha256.Sum256(etx[:])
}

type Block struct {
	Header       BlockHeader
	Transactions []Transaction
}

func (b *Block) Encode() []byte {
	block := make([]byte, 80+(len(b.Transactions)*44))
	bh := b.Header.Encode()
	copy(block[0:], bh[:])
	for k := range b.Transactions {
		tx := b.Transactions[k].Encode()
		copy(block[80+(44*k):], tx[:])
	}
	return block
}

func (b *Block) CalculateMerkle() (root [32]byte) {
	hash := sha256.New()
	for k := range b.Transactions {
		txid := b.Transactions[k].ID()
		hash.Write(txid[:])
	}
	copy(root[:], hash.Sum(nil))
	return
}
