package main

import "fmt"

/*
	Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
	Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
	Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
	Return k.
	Custom Judge:
	The judge will test your solution with the following code:
	int[] nums = [...]; // Input array
	int[] expectedNums = [...]; // The expected answer with correct length
	int k = removeDuplicates(nums); // Calls your implementation
	assert k == expectedNums.length;
	for (int i = 0; i < k; i++) {
		assert nums[i] == expectedNums[i];
	}
	If all assertions pass, then your solution will be accepted.

Example 1:

	Input: nums = [1,1,2]
	Output: 2, nums = [1,2,_]
	Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
	It does not matter what you leave beyond the returned k (hence they are underscores).
	Example 2:

	Input: nums = [0,0,1,1,1,2,2,3,3,4]
	Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
	Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
	It does not matter what you leave beyond the returned k (hence they are underscores).

Constraints:

	1 <= nums.length <= 3 * 104
	-100 <= nums[i] <= 100
	nums is sorted in non-decreasing order.
*/
func removeDuplicatesAndReturnCount(numbers []int) []int {
	emptyIndex := -1
	for i := 0; i < len(numbers); i++ {
		var nextElement int
		if i < len(numbers)-1 {
			nextElement = numbers[i+1]
		}
		if nextElement == numbers[i] {
			if emptyIndex == -1 {
				emptyIndex = i + 1
			}
		} else {
			if emptyIndex != -1 && i < len(numbers)-1 {
				numbers[emptyIndex] = nextElement
				emptyIndex++
			}
		}
	}
	if emptyIndex == -1 {
		return numbers
	} else {
		return numbers[:emptyIndex]
	}
}

/*
	func removeDuplicates(nums []int) int {
	    j := 0
	    for i := 1; i < len(nums); i++ {
	        if nums[j] != nums[i] {
	            j++
	            nums[j] = nums[i]
	        }
	    }
	    nums = nums[:j+1]
	    return len(nums)
	}
*/
func main() {
	t1 := []int{1, 1, 2}
	t2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	t3 := []int{1, 1, 2, 2, 2, 3, 4, 5, 6}
	t4 := []int{}
	t5 := []int{1}
	t6 := []int{0, 0, 0, 0, 0, 0}
	t7 := []int{-3, -1, -1}
	fmt.Println(removeDuplicatesAndReturnCount(t1))
	fmt.Println(removeDuplicatesAndReturnCount(t2))
	fmt.Println(removeDuplicatesAndReturnCount(t3))
	fmt.Println(removeDuplicatesAndReturnCount(t4))
	fmt.Println(removeDuplicatesAndReturnCount(t5))
	fmt.Println(removeDuplicatesAndReturnCount(t6))
	fmt.Println(removeDuplicatesAndReturnCount(t7))

}
