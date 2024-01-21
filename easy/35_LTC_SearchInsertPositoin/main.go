package main

/*
	func searchInsert(nums []int, target int) int {
		l, r := 0, len(nums)-1

	    for l <= r {
	        mid := (l+r)/2
	        if target > nums[mid] {
	            l = mid + 1
	        } else if target < nums[mid] {
	            r = mid - 1
	        } else {
	            return mid
	        }
	    }

	    return l
	}
*/
func searchInsert(nums []int, target int) int {
	p1 := 0
	p2 := len(nums) - 1
	for p1 <= p2 {
		mid := (p1 + p2) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			p1 = mid + 1
		} else {
			p2 = mid - 1
		}
	}
	return p2 + 1
}
