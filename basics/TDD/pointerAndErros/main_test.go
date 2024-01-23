package main

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		expect := Bitcoin(10)
		assertBalance(t, wallet, expect)
	})
	t.Run("WithDraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.WithDraw(Bitcoin(10))
		expect := Bitcoin(10)
		assertNoError(t, err)
		assertBalance(t, wallet, expect)
	})
	t.Run("Insufficient funds", func(t *testing.T) {
		var initialBalance = Bitcoin(20)
		wallet := Wallet{initialBalance}
		err := wallet.WithDraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, initialBalance)
	})

}

func assertBalance(t testing.TB, wallet Wallet, expect Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != expect {
		t.Errorf("got %s expect %s", got, expect)
	}

}
func assertError(t testing.TB, err error, expect error) {
	t.Helper()
	if err == nil {
		t.Error("Not enough funds error didn't get triggered")
	}
	if !errors.Is(err, expect) {
		t.Errorf("got %q, expect %q", err, expect)
	}
}
func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Error("Not error should be trowed")
	}
}
