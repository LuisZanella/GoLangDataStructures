package main

import (
	"fmt"
)

func mergeSort(originalArray []int, helper []int, firstIndex int, lastIndex int) {
	if firstIndex < lastIndex {
		middleIndex := (firstIndex + lastIndex) / 2
		mergeSort(originalArray, helper, firstIndex, middleIndex)
		mergeSort(originalArray, helper, middleIndex+1, lastIndex)
		merge(originalArray, helper, firstIndex, middleIndex, lastIndex)
	}
}

func merge(originArray []int, helper []int, firstIndex int, middleIndex int, lastIndex int) {
	for i := firstIndex; i <= lastIndex; i++ {
		helper[i] = originArray[i]
	}
	helperLeft := firstIndex
	helperRight := middleIndex + 1
	current := firstIndex
	for helperLeft <= middleIndex && helperRight <= lastIndex {
		if helper[helperLeft] <= helper[helperRight] {
			originArray[current] = helper[helperLeft]
			helperLeft++
		} else {
			originArray[current] = helper[helperRight]
			helperRight++
		}
		current++
	}
	remaining := middleIndex - helperLeft
	for i := 0; i <= remaining; i++ {
		originArray[current+i] = helper[helperLeft+i]
	}
}
func setValues(newValue []int, originArray *[]int, helper *[]int) {
	*originArray = newValue
	*helper = make([]int, len(*originArray))
}
func main() {
	var myArray []int
	var helper []int
	setValues([]int{1, 2, 3}, &myArray, &helper)

	fmt.Println(myArray)
	mergeSort(myArray, helper, 0, len(myArray)-1)
	fmt.Println(myArray)

	setValues([]int{6, 10, 2, 50}, &myArray, &helper)

	fmt.Println(myArray)
	mergeSort(myArray, helper, 0, len(myArray)-1)
	fmt.Println(myArray)

	setValues([]int{0, 4, 6, 78, 89, 100, -100, 0}, &myArray, &helper)
	fmt.Println(myArray)
	mergeSort(myArray, helper, 0, len(myArray)-1)
	fmt.Println(myArray)

}
