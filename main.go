package main

import (
	"fmt"
)

func main() {
	bc := NewBlockchain()

	// Initial user
	alice := NewUser("Alice", 1, 1, 1998, bc)
	bob := NewUser("Bob", 3, 12, 2004, bc)
	charlie := NewUser("Charlie", 13, 11, 2007, bc)
	dave := NewUser("Dave", 26, 3, 1980, bc)
	eve := NewUser("Eve", 31, 12, 1999, bc)

	// Initial test
	alice.PerformTransaction(bob, "Test", 100)
	charlie.PerformTransaction(dave, "Test2", 320)
	eve.PerformTransaction(alice, "Test3", 50)

	// Blockchain initial state
	fmt.Println(bc)

	// Append new transaction
	bob.PerformTransaction(charlie, "Test4", 200)
	dave.PerformTransaction(eve, "Test5", 1000)

	// Modified blockchain state
	fmt.Println(bc)

	// Check if a local state is updated
	fmt.Println(alice.LocalBlockchain().Equals(bc))
}
