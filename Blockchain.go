package main

import (
	"fmt"
)

type Blockchain struct {
	blockList     []*Block
	notifiedUsers []*User
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		blockList:     []*Block{},
		notifiedUsers: []*User{},
	}
}

func (bc *Blockchain) AddUser(user *User) {
	bc.notifiedUsers = append(bc.notifiedUsers, user)
	for _, u := range bc.notifiedUsers {
		if u != user {
			u.SetLocalBlockchain(bc)
		}
	}
}

func (bc *Blockchain) AddTransaction(transaction *Transaction) {
	var block *Block
	if len(bc.blockList) == 0 {
		block = NewBlock(0, transaction)
	} else {
		block = NewBlock(bc.blockList[len(bc.blockList)-1].Hash(), transaction)
	}

	if bc.CheckNonce(transaction.Performed) {
		bc.blockList = append(bc.blockList, block)
		for _, user := range bc.notifiedUsers {
			user.SetLocalBlockchain(bc)
		}
	} else {
		fmt.Println("Transaction cannot be performed:", transaction)
	}
}

func (bc *Blockchain) CheckNonce(performed *User) bool {
	for _, user := range bc.notifiedUsers {
		if user != performed && user.LocalBlockchain() != bc {
			return false
		}
	}
	return true
}

func (bc *Blockchain) Equals(o interface{}) bool {
	other, ok := o.(*Blockchain)
	if !ok {
		return false
	}

	if len(other.blockList) != len(bc.blockList) {
		return false
	}

	for i := 1; i < len(bc.blockList); i++ {
		if bc.blockList[i].PreviousHash() != bc.blockList[i-1].Hash() {
			return false
		}
	}

	return true
}

func (bc *Blockchain) String() string {
	return fmt.Sprintf("%v", bc.blockList)
}
