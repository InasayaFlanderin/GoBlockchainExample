package main

import (
	"fmt"
	"hash/fnv"
	"time"
)

type Transaction struct {
	Performed *User
	Received  *User
	Data      string
	Date      time.Time
	Amount    int64
}

func NewTransaction(performed, received *User, data string, date time.Time, amount int64) *Transaction {
	if amount < 0 {
		panic("Amount must be positive")
	}
	return &Transaction{
		Performed: performed,
		Received:  received,
		Data:      data,
		Date:      date,
		Amount:    amount,
	}
}

func (t *Transaction) hashCode() int {
	h := fnv.New32a()
	if _, err := h.Write([]byte(fmt.Sprintf("%d", t.Performed.ID()))); err != nil {
		panic(err)
	}
	if _, err := h.Write([]byte(fmt.Sprintf("%d", t.Received.ID()))); err != nil {
		panic(err)
	}
	if _, err := h.Write([]byte(t.Data)); err != nil {
		panic(err)
	}
	if _, err := h.Write([]byte(t.Date.String())); err != nil {
		panic(err)
	}
	if _, err := h.Write([]byte(fmt.Sprintf("%d", t.Amount))); err != nil {
		panic(err)
	}
	return int(h.Sum32())
}
