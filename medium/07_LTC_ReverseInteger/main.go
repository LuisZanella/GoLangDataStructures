package main

import (
	"fmt"
	"math"
)

//**
//
//Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside
//the signed 32-bit integer range [-231, 231 - 1], then return 0.
//
//Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
//
//*/

func main() {
	fmt.Println(reverse(45))
	fmt.Println(reverse(50))
	fmt.Println(reverse(46))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	var result int
	for x != 0 {
		digit := x % 10
		x /= 10
		if result > (math.MaxInt32-digit)/10 {
			return 0
		}
		if result < (math.MinInt32-digit)/10 {
			return 0
		}
		result = result*10 + digit
	}
	return result
}
