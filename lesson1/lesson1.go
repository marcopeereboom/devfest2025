package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/marcopeereboom/devfest2025/bitcoin"
)

func wait() {
	bufio.NewReader(os.Stdin).ReadString('\n')
}

var (
	blockchain  = make(map[[32]byte]bitcoin.Block)
	genesisHash = [32]byte{} // genesis hash is predetermined
	blocks      = 5          // // number of blocks generated
)

func main() {
	// Lesson 1 - Stack em!
	previousBlockHash := genesisHash
	for i := 0; i < blocks; i++ {
		block := bitcoin.Block{
			Header: bitcoin.BlockHeader{
				PreviousBlockHash: previousBlockHash,
			},
			Transactions: []bitcoin.Transaction{},
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
	wait()

	// Lesson 1.2 - Fork blockchain
	block3, _ := hex.DecodeString("f9a267a1e76e0dd52c80124397b8aa9c4c4b80718014f0c88dc1daae868b68ec")
	copy(previousBlockHash[:], block3)
	for i := 3; /* start at height 3 */ i < 3+2; /* add 2 blocks */ i++ {
		block := bitcoin.Block{
			Header: bitcoin.BlockHeader{
				PreviousBlockHash: previousBlockHash,
				Nonce:             1, // new block hash at same height
			},
			Transactions: []bitcoin.Transaction{},
		}
		blockHash := sha256.Sum256(block.Encode())
		blockchain[blockHash] = block
		fmt.Printf("inserted block %v' - %0x\n", i+1, blockHash)
		previousBlockHash = blockHash
	}
	fmt.Printf("total blockchain length %v\n", len(blockchain))
	wait()

	// Lesson 1.3 - Show fork
	lastBlockHash = previousBlockHash
	for height := len(blockchain) - 2; !bytes.Equal(lastBlockHash[:], genesisHash[:]); height-- {
		block := blockchain[lastBlockHash]
		fmt.Printf("block %v - %0x\n", height, sha256.Sum256(block.Encode()))
		lastBlockHash = block.Header.PreviousBlockHash
	}
}
