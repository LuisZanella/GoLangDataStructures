package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("Validate empty array", func(t *testing.T) {
		got := Search([]int{}, 0)
		assertIntValues(t, got, -1)
	})
	t.Run("Validate normal case", func(t *testing.T) {
		got := Search([]int{4, 5, 1, 2, 3}, 3)
		assertIntValues(t, got, 4)
	})
	t.Run("Invalid case", func(t *testing.T) {
		got := Search([]int{4, 5, 1, 2, 3}, 8)
		assertIntValues(t, got, -1)
	})
}

func assertIntValues(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
