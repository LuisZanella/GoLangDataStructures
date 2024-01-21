package main

import (
	"slices"
	"testing"
)

func TestSumArray(t *testing.T) {
	t.Run("Positive Numbers", func(t *testing.T) {
		var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		res := SumArray(array)
		expected := 55
		if res != expected {
			t.Errorf("Received: %d but %d was expected, from %v", res, expected, array)
		}
	})
}

func TestSumTwoArraysPerPosition(t *testing.T) {
	t.Run("Arrays have different size", func(t *testing.T) {
		a1 := []int{1, 2, 3, 5}
		a2 := []int{0, 2, 3}
		res, err := SumTwoArraysPerPosition(a1, a2)
		if err == nil || len(res) > 0 {
			t.Errorf("Error wasn't triggered and function returned: %v %v", res, err)
		}
	})
	t.Run("Sums elements per position", func(t *testing.T) {
		a1 := []int{1, 2, 3, 5}
		a2 := []int{0, 2, 3, -16}
		res, _ := SumTwoArraysPerPosition(a1, a2)
		expected := []int{1, 4, 6, -11}
		if !slices.Equal(res, expected) {
			t.Errorf("Received: %v but %v was expected", res, expected)
		}
	})
}

func BenchmarkSumArray(b *testing.B) {
	b.Run("Testing bench numbers", func(b *testing.B) {
		var values []int
		for i := 0; i < b.N; i++ {
			values = append(values, i)
			SumArray(values)
		}
	})
}
