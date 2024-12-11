package main

import "fmt"

/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length)
such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums,
or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.
*/

/*
 1. Set two pointers (L & R)
 2. Verify if L or R are more than mid or than R
 3. Verify if target is more than L or R but less
 4. Repeat over L or R
*/

func Search(nums []int, target int) (indexFound int) {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
				continue
			}
			l = mid + 1
		} else {
			if nums[r] >= target && target > nums[mid] {
				l = mid + 1
				continue
			}
			r = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(Search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
