package main

import "testing"

func TestSqrt(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		got := mySqrt(0)
		want := 0
		assertion(t, got, want)
	})
	t.Run("2", func(t *testing.T) {
		got := mySqrt(2)
		want := 1
		assertion(t, got, want)
	})
	t.Run("15", func(t *testing.T) {
		got := mySqrt(15)
		want := 3
		assertion(t, got, want)
	})
	t.Run("2,147,483,647", func(t *testing.T) {
		got := mySqrt(2147483648)
		want := 46340
		assertion(t, got, want)
	})
	t.Run("2,147,483,648", func(t *testing.T) {
		got := mySqrt(2147483648)
		want := 46340
		assertion(t, got, want)
	})
}

func assertion(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("Got: %d Want: %d", got, want)
	}
}
