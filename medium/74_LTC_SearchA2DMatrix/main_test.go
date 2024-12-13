package main

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	t.Run("Normal Matrix", func(t *testing.T) {
		got := searchMatrix([][]int{
			{1, 3, 5},
			{7, 10, 11},
			{16, 20, 23},
			{30, 34, 60}}, 3)

		assertBool(t, got, true)
	})
	t.Run("Small Size", func(t *testing.T) {
		got := searchMatrix([][]int{
			{1},
			{3}}, 2)

		assertBool(t, got, false)
	})
}

func assertBool(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
