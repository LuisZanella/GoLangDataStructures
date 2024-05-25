package main

import "fmt"

func main() {
	result := ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0)
	fmt.Printf("%v", result)
	result = ThreeNumberSum([]int{}, 0)
	fmt.Printf("%v", result)
}
func ThreeNumberSum(array []int, target int) [][]int {
	var result = make([][]int, 0)
	array = mergeSort(array)
	for i, number := range array {
		leftPointer := i + 1
		rightPointer := len(array) - 1
		for leftPointer < rightPointer {
			var sumUp int = number + array[leftPointer] + array[rightPointer]
			if sumUp == target {
				result = append(result, []int{number, array[leftPointer], array[rightPointer]})
				leftPointer++
				rightPointer--
			}
			if sumUp > target {
				rightPointer--
			}
			if sumUp < target {
				leftPointer++
			}
		}

	}
	return result
}

func mergeSort(array []int) []int {
	return splitArrayAndMerge(array)
}

func splitArrayAndMerge(array []int) []int {
	if len(array) < 2 {
		return array
	}
	first := splitArrayAndMerge(array[:len(array)/2])
	second := splitArrayAndMerge(array[len(array)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	var final []int
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
