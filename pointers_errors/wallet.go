/**
We have covered:
- struct pointers are automatically dereferenced.
- nil: the `zero value` for pointer, func, error, etc.,
- check errors and handle them gracefully
- create new type let you implement interfaces
*/

package main

import (
	"errors"
	"fmt"
)

// Adding `var` makes it global in the package
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	// In Go if a symbol starts with a lowercase symbol then it is PRIVATE
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}