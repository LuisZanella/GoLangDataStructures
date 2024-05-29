package main

import "fmt"

func main() {
	res := MaxSubsetSumNoAdjacent([]int{75, 105, 120, 75, 90, 135})
	fmt.Printf("%v", res)
}
func MaxSubsetSumNoAdjacent(array []int) int {
	var sum int
	var adjacent bool
	for i := 0; i < len(array); i++ {
		if i != 0 || i != len(array)-1 || i == len(array)-2 {
			if adjacent && i != len(array)-2 {
				adjacent = false
				continue
			}
		}
		sum += array[i]
		adjacent = true
	}
	return sum
}
