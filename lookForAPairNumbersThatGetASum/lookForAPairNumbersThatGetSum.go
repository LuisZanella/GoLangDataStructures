package main

import "fmt"

func main() {
	array1 := []int{1, 2, 3, 9} // no
	array2 := []int{1, 2, 4, 4} // yes
	// array2 := []int{1, 2, 6, 4} // yes
	var totalToSum int = 8
	fmt.Println(hasPairEqualToSum(array1, totalToSum))
	fmt.Println(hasPairEqualToSumEfficent(array2, totalToSum))

}

func hasPairEqualToSum(array []int, total int) bool {
	for i := 0; i < len(array); i++ {
		for ii := 0; ii < len(array); ii++ {
			if (array[i] + array[ii]) == total {
				return true
			}
		}
	}
	return false
} // O(n^2)
func hasPairEqualToSumEfficent(array []int, total int) bool {
	var complementNumbers []int
	for _, number := range array {
		if contains(complementNumbers, number) {
			return true
		}
		complementNumbers = append(complementNumbers, (total - number))
	}
	return false
} // O (n)
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
} // O(n)
