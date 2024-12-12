package main

import "fmt"

// log all pairs of array
var array = [5]string{"a", "b", "c", "d", "e"}

func main() {
	BIG_O_NxN()
	BIG_O_N()
	BIG_O_1()
}

func BIG_O_NxN() { // BIG-O O(n * n) => O(n^2)
	for index1, letter := range array {
		for index2, letter2 := range array {
			fmt.Printf("%s %s %d %d \n", letter, letter2, index1, index2)
		}
	}
}
func BIG_O_N() { // BIG-O O(n)
	for index1, letter := range array {
		fmt.Printf("%s %d \n", letter, index1)
	}
}
func BIG_O_1() { // BIG-O O(1)
	fmt.Printf("%s \n", array[0])
}
