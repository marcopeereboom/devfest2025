package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/bits"
	"time"

	"github.com/marcopeereboom/devfest2025/bitcoin"
)

func main() {
	miner := [20]byte{'m', 'i', 'n', 'e', 'r'}
	alice := [20]byte{'a', 'l', 'i', 'c', 'e'}
	transaction := bitcoin.Transaction{
		From:   miner,
		To:     alice,
		Amount: 50,
	}

	// Lesson 3 - Mining
	var genesisHash [32]byte
	copy(genesisHash[:], bytes.Repeat([]byte{'%'}, 32))
	copy(genesisHash[:], []byte("genesishash"))
	block := bitcoin.Block{
		Header: bitcoin.BlockHeader{
			Version:           1,
			PreviousBlockHash: genesisHash,
			MerkleRoot:        [32]byte{},
			Timestamp:         uint32(time.Now().Unix()),
			Difficulty:        3 * 8, // 24 bits of leading zeroes
			Nonce:             0,     // This is the key to mining!
		},
		Transactions: []bitcoin.Transaction{
			transaction,
		},
	}
	block.Header.MerkleRoot = block.CalculateMerkle()

	var blockHash [32]byte
	start := time.Now()
	for i := uint32(0); i < math.MaxUint32; i++ {
		block.Header.Nonce = i // It is this simple!
		encodedBlockHeader := block.Header.Encode()
		blockHash = sha256.Sum256(encodedBlockHeader[:])
		x := binary.BigEndian.Uint64(blockHash[0:8])
		if bits.LeadingZeros64(x) >= int(block.Header.Difficulty) {
			break
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("mining   : hashes %v time %v hashes/second %f\n", block.Header.Nonce,
		elapsed, float64(block.Header.Nonce)/float64(elapsed.Seconds()))
	fmt.Printf("nonce    : 0x%0x\n", block.Header.Nonce)
	fmt.Printf("merkle   : %0x\n", block.Header.MerkleRoot)
	fmt.Printf("blockhash: %0x\n", blockHash[:])
	fmt.Printf("block    :\n%v\n", hex.Dump(block.Encode()))
}
