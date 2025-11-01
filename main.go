package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/bits"
	"os"
	"time"
)

func lesson1() {
	blockchain := make(map[[32]byte]Block)
	genesisHash := [32]byte{} // genesis hash is predetermined

	// Lesson 1 - Stack em!
	blocks := 5
	previousBlockHash := genesisHash
	for i := 0; i < blocks; i++ {
		block := Block{
			Header: BlockHeader{
				PreviousBlockHash: previousBlockHash,
			},
			Transactions: []Transaction{},
		}
		blockHash := sha256.Sum256(block.Encode())
		blockchain[blockHash] = block
		fmt.Printf("inserted block %v - %0x\n", i, blockHash)
		previousBlockHash = blockHash
	}
	wait()

	// Lesson 1.1 - Print blockchain
	lastBlockHash := previousBlockHash
	for height := blocks - 1; !bytes.Equal(lastBlockHash[:], genesisHash[:]); height-- {
		block := blockchain[lastBlockHash]
		fmt.Printf("block %v - %0x\n", height, sha256.Sum256(block.Encode()))
		lastBlockHash = block.Header.PreviousBlockHash
	}

	// Lesson 1.2 - Fork blockchain
	block3, _ := hex.DecodeString("f9a267a1e76e0dd52c80124397b8aa9c4c4b80718014f0c88dc1daae868b68ec")
	copy(previousBlockHash[:], block3)
	for i := 3; /* start at height 3 */ i < 3+2; /* add 2 blocks */ i++ {
		block := Block{
			Header: BlockHeader{
				PreviousBlockHash: previousBlockHash,
				Nonce:             1, // new block hash at same height
			},
			Transactions: []Transaction{},
		}
		blockHash := sha256.Sum256(block.Encode())
		blockchain[blockHash] = block
		fmt.Printf("inserted block %v' - %0x\n", i, blockHash)
		previousBlockHash = blockHash
	}
	fmt.Printf("total blockchain length %v\n", len(blockchain))

	// Lesson 1.3 - Show fork
	lastBlockHash = previousBlockHash
	for height := len(blockchain) - 2; !bytes.Equal(lastBlockHash[:], genesisHash[:]); height-- {
		block := blockchain[lastBlockHash]
		fmt.Printf("block %v - %0x\n", height, sha256.Sum256(block.Encode()))
		lastBlockHash = block.Header.PreviousBlockHash
	}
}

func lesson2() {
	alice := [20]byte{'a', 'l', 'i', 'c', 'e'}
	bob := [20]byte{'b', 'o', 'b'}

	// Lesson 2 - Transaction
	transaction := Transaction{
		From:   alice,
		To:     bob,
		Amount: 50,
	}
	tx := transaction.Encode()
	fmt.Printf("txid   : %0x\n", sha256.Sum256(tx[:]))
	fmt.Printf("payload:\n%v\n", hex.Dump(tx[:]))
}

func lesson3() {
	miner := [20]byte{'m', 'i', 'n', 'e', 'r'}
	alice := [20]byte{'a', 'l', 'i', 'c', 'e'}
	transaction := Transaction{
		From:   miner,
		To:     alice,
		Amount: 50,
	}

	// Lesson 3 - Mining
	var genesisHash [32]byte
	copy(genesisHash[:], bytes.Repeat([]byte{'%'}, 32))
	copy(genesisHash[:], []byte("genesishash"))
	block := Block{
		Header: BlockHeader{
			Version:           1,
			PreviousBlockHash: genesisHash,
			MerkleRoot:        [32]byte{},
			Timestamp:         uint32(time.Now().Unix()),
			Difficulty:        3 * 8, // 24 bits of leading zeroes
			Nonce:             0,
		},
		Transactions: []Transaction{
			transaction,
		},
	}
	block.Header.MerkleRoot = block.CalculateMerkle()

	var blockHash [32]byte
	start := time.Now()
	for i := uint32(0); i < math.MaxUint32; i++ {
		block.Header.Nonce = i
		encodedBlockHeader := block.Header.Encode()
		blockHash = sha256.Sum256(encodedBlockHeader[:])
		x := binary.BigEndian.Uint64(blockHash[0:8])
		// XXX Replace endian with hand rolled encoders?
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

func newBlock(previousBlockHash [32]byte) Block {
	return Block{
		Header: BlockHeader{
			Version:           1,
			PreviousBlockHash: previousBlockHash,
			MerkleRoot:        [32]byte{},
			Timestamp:         uint32(time.Now().Unix()),
			Difficulty:        3 * 8, // 24 bits of leading zeroes
			Nonce:             0,
		},
	}
}

func mineBlock(block *Block, rewardAddress [20]byte, transactions []Transaction) (*[32]byte, []byte) {
	miner := [20]byte{'m', 'i', 'n', 'e', 'r'}
	block.Transactions = append([]Transaction{{
		From:   miner,
		To:     rewardAddress,
		Amount: 50,
	}}, transactions...)
	block.Header.MerkleRoot = block.CalculateMerkle()

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

func broadcast(blockchain map[[32]byte]Block, addresses map[[20]byte]uint32, block *Block) {
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

func lesson4() {
	alice := [20]byte{'a', 'l', 'i', 'c', 'e'}
	bob := [20]byte{'b', 'o', 'b'}

	var genesisHash [32]byte
	copy(genesisHash[:], bytes.Repeat([]byte{'%'}, 32))
	copy(genesisHash[:], []byte("genesishash"))

	blockchain := make(map[[32]byte]Block)
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
		blockHash, _ := mineBlock(&block, alice, []Transaction{
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

func _main() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("must provide lesson")
	}

	switch os.Args[1] {
	case "lesson1":
		lesson1()
	}

	return nil
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
