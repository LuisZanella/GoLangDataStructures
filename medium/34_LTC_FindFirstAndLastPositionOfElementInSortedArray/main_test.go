package main

import (
	"testing"
)

func TestSearchRange(t *testing.T) {
	t.Run("Empty Array", func(t *testing.T) {
		got := searchFirstOrLastIndexInRange([]int{}, 0, false)
		assertIntArray(t, got, -1)
	})
	t.Run("Target in half first index", func(t *testing.T) {
		got := searchFirstOrLastIndexInRange([]int{5, 7, 7, 8, 8, 10}, 10, false)
		assertIntArray(t, got, 5)
	})
	t.Run("Target in second half first index", func(t *testing.T) {
		got := searchFirstOrLastIndexInRange([]int{5, 7, 7, 8, 8, 10}, 5, false)
		assertIntArray(t, got, 0)
	})
	t.Run("Target returns last index", func(t *testing.T) {
		got := searchFirstOrLastIndexInRange([]int{5, 7, 7, 8, 8, 10}, 7, true)
		assertIntArray(t, got, 2)
	})
}

func assertIntArray(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
	//if !slices.Equal(got, want) {
	//	t.Errorf("got %v want %v", got, want)
	//}
}
