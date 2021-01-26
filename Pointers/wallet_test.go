package Pointers_test

import (
	"testing"
	"tests/Pointers"
)

func TestWallet(t *testing.T){

	t.Run("Deposit", func(t *testing.T) {
		wallet := &Pointers.Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := 10

		if got != want {
			t.Errorf("got %d want %d",got,want)
		}
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Pointers.NewWallet(20)

		wallet.Withdraw(10)

		got := wallet.Balance()

		want := 10

		if got != want{
			t.Errorf("got %d want %d",got,want)
		}
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Pointers.NewWallet(20)

		err := wallet.Withdraw(100)

		got := wallet.Balance()

		want := 20

		if err == nil{
			t.Errorf("wanted an error but didn't get one")
		}

		if got != want{
			t.Errorf("got %d want %d",got,want)
		}
	})
}
