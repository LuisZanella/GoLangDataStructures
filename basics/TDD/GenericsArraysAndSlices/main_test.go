package main

import (
	"strings"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("Multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})
	t.Run("Concatenate bytes", func(t *testing.T) {
		concatenate := func(x, y byte) byte {
			return x + y
		}
		AssertEqual(t, Reduce([]byte{'a', 'b', 'c'}, concatenate, 0), 38)
	})
	t.Run("Concatenate string", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}
		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}
func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)
}
func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})
	type Person struct {
		Name string
	}

	t.Run("Find the best Dev and IT lover", func(t *testing.T) {
		people := []Person{
			Person{Name: "L1nux T0rvalds"},
			Person{Name: "St3ve W0zniak"},
			Person{Name: "K3n Th0mpson"},
			Person{Name: "K3vinazo M1tnick"},
			Person{Name: "Capta1n Crush"},
			Person{Name: "Luis Zanella"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Luis")
		})

		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "Luis Zanella"})
	})
}
func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}
