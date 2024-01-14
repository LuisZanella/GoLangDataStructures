package main

import (
	"fmt"
	"time"
)

func isNumberPalindrome_NotOptimal(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var clone int = x
	var invertedNumber = 0
	for clone > 0 {
		var lastDigit = clone % 10
		invertedNumber = invertedNumber*10 + lastDigit
		clone = clone / 10
	}
	return invertedNumber == x
}

func getPow10(length int) int {
	if length <= 0 {
		return 1
	}
	decimalAdder := 10
	for i := 1; i < length; i++ {
		decimalAdder *= 10
	}
	return decimalAdder
}

func getDecimalByPosition(number int, position int) int {
	return (number / getPow10(position-1)) % getPow10(1)
}
func isNumberPalindrome_OptimalTime(number int) bool {
	if number < 0 {
		return false
	}

	var numberLength int = 0
	for clone := number; clone > 0; clone = clone / 10 {
		numberLength++
	}
	for i := 1; i <= numberLength/2; i++ {
		if getDecimalByPosition(number, i) != getDecimalByPosition(number, numberLength-i+1) {
			return false
		}
	}
	return true
}

func finishCounter(startTime int64) {
	endTime := time.Now().UnixNano()
	duration := (endTime - startTime)

	fmt.Printf("Time taken: %d unix\n", duration)
}

func main() {
	startTime := time.Now().UnixNano()
	fmt.Println(isNumberPalindrome_NotOptimal(10101))
	fmt.Println(isNumberPalindrome_NotOptimal(1))
	fmt.Println(isNumberPalindrome_NotOptimal(3213211))
	fmt.Println(isNumberPalindrome_NotOptimal(323232))
	finishCounter(startTime)

	startTime = time.Now().UnixNano()
	fmt.Println(isNumberPalindrome_OptimalTime(10101))
	fmt.Println(isNumberPalindrome_OptimalTime(1))
	fmt.Println(isNumberPalindrome_OptimalTime(3213211))
	fmt.Println(isNumberPalindrome_OptimalTime(323232))
	finishCounter(startTime)

}
