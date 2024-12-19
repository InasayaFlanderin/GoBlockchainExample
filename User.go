package main

import (
	"time"
)

type User struct {
	name            string
	birthDate       time.Time
	localBlockchain *Blockchain
	id              int64
}

func NewUser(name string, day, month, year int, local *Blockchain) *User {
	createDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	user := &User{
		name:            name,
		birthDate:       createDate,
		localBlockchain: local,
		id:              (time.Now().UnixNano() & int64(hash(name))) | (createDate.UnixNano()^32)>>3,
	}
	local.AddUser(user)
	return user
}

func (u *User) Name() string {
	return u.name
}

func (u *User) BirthDate() time.Time {
	return u.birthDate
}

func (u *User) LocalBlockchain() *Blockchain {
	return u.localBlockchain
}

func (u *User) SetLocalBlockchain(local *Blockchain) {
	u.localBlockchain = local
}

func (u *User) ID() int64 {
	return u.id
}

func (u *User) PerformTransaction(received *User, data string, amount int64) {
	if amount < 0 {
		panic("Amount must be positive")
	}
	transaction := NewTransaction(u, received, data, time.Now(), amount)
	u.localBlockchain.AddTransaction(transaction)
}

func (u *User) Equals(o interface{}) bool {
	if other, ok := o.(*User); ok {
		return other.ID() == u.id
	}
	return false
}

func hash(s string) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = 31*h + int(s[i])
	}
	return h
}
