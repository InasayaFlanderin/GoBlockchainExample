package main

import "fmt"

type Block struct {
	previousHash int
	hash         int
	transaction  *Transaction
}

func NewBlock(previousHash int, transaction *Transaction) *Block {
	block := &Block{
		previousHash: previousHash,
		transaction:  transaction,
		hash:         calculateHash(transaction, previousHash),
	}
	return block
}

func (b *Block) PreviousHash() int {
	return b.previousHash
}

func (b *Block) Hash() int {
	return b.hash
}

func (b *Block) Transaction() *Transaction {
	return b.transaction
}

func calculateHash(transaction *Transaction, previousHash int) int {
	rawHash := transaction.hashCode()
	rawHash |= 3
	rawHash ^= previousHash
	rawHash <<= 3*transaction.hashCode() + previousHash
	rawHash -= 17
	rawHash *= 31

	return rawHash
}

func (b *Block) String() string {
	return fmt.Sprintf("Block{previousHash=%d, hash=%d, transaction=%d -> %d: %s amount: %f}",
		b.previousHash, b.hash, b.transaction.Performed.ID(), b.transaction.Received.ID(), b.transaction.Data, b.transaction.Amount)
}
