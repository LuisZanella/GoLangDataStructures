package main

import "fmt"

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) -1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] > target {
			r = m - 1
		} else {
			if nums[m] < target {
				l = m + 1
			} else {
				return m
			}
		}
	}
	return l
}

func main() {
	numbers := []int{-14, 32, 20, 19}

	target := -15

	fmt.Println(searchInsert(numbers, target))
}
