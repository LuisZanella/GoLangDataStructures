package main

import "fmt"

func getLengthDifferentInArray(nums []int, val int) int {
	ei := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ei] = nums[i]
			ei++
		}

	}
	return len(nums[:ei])
}

/*
	func removeElement(nums []int, val int) int {
	    k := len(nums)
	    if k > 0 {
	        for i := k-1; i >= 0; i-- {
	            if nums[i] == val{
	                nums[k-1],nums[i] = nums[i],nums[k-1]
	                k--
	            }
	        }
	    }
	    return k
	}
*/
func main() {
	fmt.Println(getLengthDifferentInArray([]int{0, 0, 0, 0, 0, 0}, 0))
	fmt.Println(getLengthDifferentInArray([]int{1, 2, 3, 4, -2, 5}, -2))
	fmt.Println(getLengthDifferentInArray([]int{}, 0))
	fmt.Println(getLengthDifferentInArray([]int{4, 2, 1, 25}, -1))
}
