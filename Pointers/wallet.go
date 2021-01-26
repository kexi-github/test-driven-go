package Pointers

import "errors"

type Wallet struct {
	balance int
}

func NewWallet(balance int) *Wallet {
	return &Wallet{balance: balance}
}

func (w *Wallet)Deposit(c int) {
	w.balance += c
}

func (w *Wallet)Balance() int {
	return w.balance
}

func (w *Wallet)Withdraw(c int) error{

	if w.balance < c{
		return errors.New("oh no")
	}else {
		w.balance -= c
	}
	return nil
}
