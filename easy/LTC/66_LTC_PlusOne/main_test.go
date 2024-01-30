package main

import (
	"slices"
	"testing"
)

func TestPlusOne(t *testing.T) {
	var nums []int
	t.Run("3 great range", func(t *testing.T) {
		nums = []int{5, 1, 2}
		expect := []int{5, 1, 3}
		got := plusOne(nums)
		if !slices.Equal(got, expect) {
			t.Errorf("Received: %v  Expected: %v", got, expect)
		}
	})
	t.Run("3 left increment", func(t *testing.T) {
		nums = []int{5, 9, 9}
		expect := []int{6, 0, 0}
		got := plusOne(nums)
		if !slices.Equal(got, expect) {
			t.Errorf("Received: %v  Expected: %v", got, expect)
		}
	})
	t.Run("3 first increment", func(t *testing.T) {
		nums = []int{9, 9, 9}
		expect := []int{1, 0, 0, 0}
		got := plusOne(nums)
		if !slices.Equal(got, expect) {
			t.Errorf("Received: %v  Expected: %v", got, expect)
		}
	})
	t.Run("Empty", func(t *testing.T) {
		nums = []int{}
		expect := []int{1}
		got := plusOne(nums)
		if !slices.Equal(got, expect) {
			t.Errorf("Received: %v  Expected: %v", got, expect)
		}
	})
}
