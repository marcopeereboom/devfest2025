package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/bits"
	"time"

	"github.com/marcopeereboom/devfest2025/bitcoin"
)

func newBlock(previousBlockHash [32]byte) bitcoin.Block {
	return bitcoin.Block{
		Header: bitcoin.BlockHeader{
			Version:           1,
			PreviousBlockHash: previousBlockHash,
			MerkleRoot:        [32]byte{},
			Timestamp:         uint32(time.Now().Unix()),
			Difficulty:        3 * 8, // 24 bits of leading zeroes
			Nonce:             0,
		},
	}
}

func mineBlock(block *bitcoin.Block, rewardAddress [20]byte, transactions []bitcoin.Transaction) (*[32]byte, []byte) {
	miner := [20]byte{'m', 'i', 'n', 'e', 'r'}
	block.Transactions = append([]bitcoin.Transaction{{
		From:   miner,
		To:     rewardAddress,
		Amount: 50,
	}}, transactions...)
	block.Header.MerkleRoot = block.CalculateMerkle() // This links header and block
	var blockHash [32]byte
	for i := uint32(0); i < math.MaxUint32; i++ {
		block.Header.Nonce = i
		encodedBlockHeader := block.Header.Encode()
		blockHash = sha256.Sum256(encodedBlockHeader[:])
		x := binary.BigEndian.Uint64(blockHash[0:8])
		if bits.LeadingZeros64(x) >= int(block.Header.Difficulty) {
			return &blockHash, block.Encode()
		}
	}
	return nil, nil
}

func broadcast(blockchain map[[32]byte]bitcoin.Block, addresses map[[20]byte]uint32, block *bitcoin.Block) {
	encodedBlockHeader := block.Header.Encode()
	blockHash := sha256.Sum256(encodedBlockHeader[:])
	blockchain[blockHash] = *block

	for k, tx := range block.Transactions {
		if k != 0 {
			addresses[tx.From] -= tx.Amount
		}
		addresses[tx.To] += tx.Amount
	}
}

func main() {
	alice := [20]byte{'a', 'l', 'i', 'c', 'e'}
	bob := [20]byte{'b', 'o', 'b'}

	var genesisHash [32]byte
	copy(genesisHash[:], bytes.Repeat([]byte{'%'}, 32))
	copy(genesisHash[:], []byte("genesishash"))

	blockchain := make(map[[32]byte]bitcoin.Block)
	addresses := make(map[[20]byte]uint32)

	// Lesson 4 - putting it all together
	blocks := 5
	previousBlockHash := genesisHash
	for i := 0; i < blocks; i++ {
		block := newBlock(previousBlockHash)
		blockHash, _ := mineBlock(&block, alice, nil)
		broadcast(blockchain, addresses, &block)
		fmt.Printf("mined block %v - %x alice balance %v\n", i, *blockHash,
			addresses[alice])
		previousBlockHash = *blockHash
	}

	// Lesson 4.1 Alice sends Bob funds
	for i := blocks; i < blocks*2; i++ {
		block := newBlock(previousBlockHash)
		blockHash, _ := mineBlock(&block, alice, []bitcoin.Transaction{
			{
				From:   alice,
				To:     bob,
				Amount: 20,
			},
		})
		broadcast(blockchain, addresses, &block)
		fmt.Printf("mined block %v - %x alice balance %v bob balance %v\n", i,
			*blockHash, addresses[alice], addresses[bob])
		previousBlockHash = *blockHash
	}
}
