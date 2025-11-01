package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/marcopeereboom/devfest2025/bitcoin"
)

func main() {
	alice := [20]byte{'a', 'l', 'i', 'c', 'e'}
	bob := [20]byte{'b', 'o', 'b'}

	// Lesson 2 - Transaction
	transaction := bitcoin.Transaction{
		From:   alice,
		To:     bob,
		Amount: 50,
	}
	tx := transaction.Encode()
	fmt.Printf("txid   : %0x\n", sha256.Sum256(tx[:]))
	fmt.Printf("payload:\n%v\n", hex.Dump(tx[:]))
}
