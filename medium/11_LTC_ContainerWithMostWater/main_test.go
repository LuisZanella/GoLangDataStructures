package main

import "testing"

func TestGetContainerWithMostWater(t *testing.T) {
	t.Run("Empty Array", func(t *testing.T) {
		got := GetContainerWithMostWater([]int{})
		AssertPrint(t, got, 0)
	})
	t.Run("One digit", func(t *testing.T) {
		got := GetContainerWithMostWater([]int{1})
		AssertPrint(t, got, 0)
	})
	t.Run("Two digits", func(t *testing.T) {
		got := GetContainerWithMostWater([]int{1, 2})
		AssertPrint(t, got, 1)
	})
	t.Run("A normal volume", func(t *testing.T) {
		got := GetContainerWithMostWater([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
		AssertPrint(t, got, 49)
	})
}

func AssertPrint(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}
