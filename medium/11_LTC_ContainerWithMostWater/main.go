package main

//11. Container With Most Water
//Medium
//Topics
//Companies
//Hint
//You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the
//ith line are (i, 0) and (i, height[i]).
//
//Find two lines that together with the x-axis form a container, such that the container contains the most water.
//
//Return the maximum amount of water a container can store.
//
//Notice that you may not slant the container.

func main() {

}
func GetContainerWithMostWater(height []int) int {
	maxCapacity := 0
	lenHeight := len(height)
	if lenHeight < 2 {
		return maxCapacity
	}
	left := 0
	right := lenHeight - 1

	for left < right {
		sliceHeight := min(height[left], height[right])
		distance := right - left

		capacity := sliceHeight * distance
		if maxCapacity < capacity {
			maxCapacity = capacity
		}
		if sliceHeight == height[left] {
			left++
		} else {
			right--
		}
	}
	return maxCapacity
}
