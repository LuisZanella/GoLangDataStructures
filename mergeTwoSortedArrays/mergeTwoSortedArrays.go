package main

import f "fmt"

func main() {
	// array1 := []int{2, 4, 5, 6, 7}
	// array2 := []int{1, 2, 4, 5, 7, 8, 8, 9}
	array1 := []int{0, 3, 4, 31}
	array2 := []int{3, 4, 6, 30}
	f.Println(MergeAndSortArray(array1, array2))
	MergeArraysFromOtherInterfaces()
}
func MergeAndSortArray(array1 []int, array2 []int) (mergedArray []int) {
	if len(array1) == 0 {
		return array2
	}
	if len(array2) == 0 {
		return array1
	}

	var itemArray1 int
	var itemArray2 int
	var fi int = 0
	var si int = 0

	for fi < len(array1) || si < len(array2) {
		if fi < len(array1) {
			itemArray1 = array1[fi]
		}
		if si < len(array2) {
			itemArray2 = array2[si]
		}
		if itemArray1 < itemArray2 && fi < len(array1) || si == len(array2) {
			mergedArray = append(mergedArray, itemArray1)
			fi++
		} else {
			if itemArray2 < itemArray1 && si < len(array2) || fi == len(array1) {
				mergedArray = append(mergedArray, itemArray2)
				si++
			}
		}
		if itemArray1 == itemArray2 {
			mergedArray = append(mergedArray, itemArray1)
			mergedArray = append(mergedArray, itemArray2)
			fi++
			si++
		}

	}
	return mergedArray

}
func MergeArraysFromOtherInterfaces() {
	array2 := [...]string{"wa", "2ad"}
	array1 := [...]int{1, 2, 3, 4}

	interfaceArray := make([]interface{}, len(array1))
	interfaceArray2 := make([]interface{}, len(array2))

	for index, data := range array1 {
		interfaceArray[index] = data
	}

	for index, data := range array2 {
		interfaceArray2[index] = data
	}
	f.Println(mergeArrays(interfaceArray, interfaceArray2))
}
func mergeArrays(array1 []interface{}, array2 []interface{}) (mergedArray []interface{}) {
	mergedArray = append(array1, array2...)
	return mergedArray
}

// Doesn't Work
// package main

// import (
// 	f "fmt"
// )

// func main() {
// 	array2 := []string{"wa", "2ad"}
// 	array1 := []int{1, 2, 3, 4}
// 	f.Println(mergeArrays(array2, array1))
// }
// func mergeArrays(array1 interface{}, array2 interface{}) (mergedArray []interface{}) {
// 	mergedArray = AppendInterfaceToInterfaceArray(array1, mergedArray)
// 	f.Println(mergedArray)
// 	mergedArray = AppendInterfaceToInterfaceArray(array2, mergedArray)
// 	f.Println(mergedArray)
// 	return mergedArray
// }
// func AppendInterfaceToInterfaceArray(array interface{}, actualArray []interface{}) (mergedArray []interface{}) {
// 	switch array := array.(type) {
// 	case []string:
// 		for _, value := range array {
// 			mergedArray = append([]interface{}{actualArray}, value)
// 			f.Println(mergedArray)
// 		}
// 	case []int:
// 		for _, value := range array {
// 			mergedArray = append([]interface{}{actualArray}, value)
// 			f.Println(mergedArray)
// 		}
// 	default:
// 		f.Println(array)
// 	}
// 	return mergedArray
// }
