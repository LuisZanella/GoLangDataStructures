package main

import "fmt"

//[a,b,c,d] [f,g,s,d] = d => d yes
//[a,b,c,r] [f,g,s,d] no
// i can use all memory that i want
// brutal solution => array compared with the other array for each element

func main() {
	array1 := []byte{'a', 'b', 'c', 'x'}
	array2 := []byte{'z', 'y', 'i', 'r'}
	array3 := []byte{'z', 'y', 'x'}
	fmt.Println(arraysContainsCommonItems(array1, array2))
	fmt.Println(arraysContainsCommonItems(array1, array3))
	fmt.Println(arraysContainsCommonItems(array2, array3))
}

func arraysContainsCommonItems(array1 []byte, array2 []byte) bool {
	//loop through first array and create object
	//where properties === item in the array
	tableHashArray1 := map[string]bool{}
	for _, ele := range array1 {
		if !tableHashArray1[string(ele)] {
			tableHashArray1[string(ele)] = true
		}
	}
	//loop through array and check if item in second
	//array  exist on created object
	for _, ele := range array2 {
		if tableHashArray1[string(ele)] {
			return true
		}
	}
	return false
} // O(a+b) Time complexity :)
