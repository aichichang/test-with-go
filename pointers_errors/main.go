package main

import (
	"errors"
	"fmt"
)

// Pointers:
// Go copies values when you pass them to functions/methods,
// so if you're writing a function that needs to mutate state
// you'll need it to take a pointer to the thing you want to change.

// References:
// sometimes you won't want your system to make a copy of something, in which case you need to pass a reference.
// Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).

// Any struct that has a String method is part of the Stringer interface
type Stringer interface {
	String() string
}

type Bitcoin int

// Turn Bitcoin into a formatted string
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// errors.New creates a new error with a message of your choosing
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// pointer to the real Wallet, so we can add value to it
func (w *Wallet) Deposite(amount Bitcoin) {
	w.balance += amount
}

// We have to return the Bitcoin type due to w.balance's type
func (w *Wallet) Balance() Bitcoin {
	//struct pointers and they are automatically dereferenced
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
