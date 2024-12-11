package main

import "fmt"

/*
Given an array of integers nums sorted in non-decreasing order,
find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.
*/
func searchRange(nums []int, target int) (indexRange []int) {
	start := searchFirstOrLastIndexInRange(nums, target, false)
	if start == -1 {
		return []int{-1, -1}
	}
	last := searchFirstOrLastIndexInRange(nums, target, true)
	return []int{start, last}
}

func searchFirstOrLastIndexInRange(nums []int, target int, last bool) (indexFound int) {
	l := 0
	r := len(nums) - 1
	indexFound = -1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			indexFound = m
			if last {
				l = m + 1
			} else {
				r = m - 1
			}
			continue
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return indexFound
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}
