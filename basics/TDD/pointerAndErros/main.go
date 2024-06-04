package main

import (
	"errors"
	"fmt"
)

type Bitcoin int
type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("You need more bitcoins!!!<3")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}
func (w *Wallet) WithDraw(value Bitcoin) error {
	if value > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= value
	return nil
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//TEST
