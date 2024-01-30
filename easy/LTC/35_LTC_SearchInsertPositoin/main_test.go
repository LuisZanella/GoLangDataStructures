package main

import "testing"

func TestSearchInsert(t *testing.T) {
	var numbers []int
	var target int
	validateOutput := func(value []int, got int, expect int) {
		if got != expect {
			t.Errorf("%v array got: %d but was expecting %d", value, got, expect)
		}
	}
	t.Run("Empty", func(t *testing.T) {
		numbers = []int{}
		target = 3
		got := searchInsert(numbers, target)
		expect := 0
		validateOutput(numbers, got, expect)
	})
	t.Run("Negative", func(t *testing.T) {
		numbers = []int{-14, 32, 20, 19}
		target = -15
		got := searchInsert(numbers, target)
		expect := 0
		validateOutput(numbers, got, expect)
	})
	t.Run("Positive", func(t *testing.T) {
		numbers = []int{1, 3, 4, 5, 6, 7, 8, 9}
		target = 2
		got := searchInsert(numbers, target)
		expect := 1
		validateOutput(numbers, got, expect)
	})
	t.Run("Middle", func(t *testing.T) {
		numbers = []int{1, 2, 3, 4, 6, 7, 8, 9}
		target = 5
		got := searchInsert(numbers, target)
		expect := 4
		validateOutput(numbers, got, expect)
	})
	t.Run("Higher", func(t *testing.T) {
		numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		target = 10
		got := searchInsert(numbers, target)
		expect := 9
		validateOutput(numbers, got, expect)
	})
}
