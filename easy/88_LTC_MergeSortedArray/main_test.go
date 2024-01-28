package main

import (
	"reflect"
	"testing"
)

func assert(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got: %v Want: %v", got, want)
	}
}
func TestMerge(t *testing.T) {
	var a1, a2 []int
	var m, n int
	t.Run("nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3", func(t *testing.T) {
		a1 = []int{1, 2, 3, 0, 0, 0}
		a2 = []int{2, 5, 6}
		m = 3
		n = 3
		merge(a1, m, a2, n)
		want := []int{1, 2, 2, 3, 5, 6}
		assert(t, a1, want)
	})
	t.Run("nums1 = [1], m = 1, nums2 = [], n = 0", func(t *testing.T) {
		a1 = []int{1}
		a2 = []int{}
		m = 1
		n = 0
		merge(a1, m, a2, n)
		want := []int{1}
		assert(t, a1, want)
	})
	t.Run("nums1 = [0], m = 0, nums2 = [1], n = 1", func(t *testing.T) {
		a1 = []int{0}
		a2 = []int{1}
		m = 0
		n = 1
		merge(a1, m, a2, n)
		want := []int{1}
		assert(t, a1, want)
	})

}
