package main

import "testing"

func TestClimbStairs(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		got := climbStairs(0)
		want := 1
		AssertPrint(t, got, want)
	})
	t.Run("1", func(t *testing.T) {

		got := climbStairs(1)
		want := 1
		AssertPrint(t, got, want)
	})
	t.Run("2", func(t *testing.T) {
		got := climbStairs(2)

		want := 2
		AssertPrint(t, got, want)
	})
	t.Run("3", func(t *testing.T) {

		got := climbStairs(3)
		want := 3
		AssertPrint(t, got, want)
	})
	t.Run("45", func(t *testing.T) {

		got := climbStairs(45)
		want := 1134903170
		AssertPrint(t, got, want)
	})
	//t.Run("50", func(t *testing.T) {
	//
	//	got := climbStairs(50)
	//	want := 12586269025
	//	AssertPrint(t, got, want)
	//})
	//t.Run("100", func(t *testing.T) {
	//
	//	got := climbStairs(100)
	//	want := 354224848179261915075
	//	AssertPrint(t, got, want)
	//})
}

func AssertPrint(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %d Want: %d", got, want)
	}
}
