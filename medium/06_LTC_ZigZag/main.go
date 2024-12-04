package main

import (
	"fmt"
	"strings"
)

// **
// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to
// display this pattern in a fixed font for better legibility)
//
// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
//
// Write the code that will take a string and make this conversion given a number of rows:
//
// string convert(string s, int numRows);
//
// */
func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert2("PAYPALISHIRING", 3))
	fmt.Println(convert2("AB", 1))
}


func convert(s string, numRows int) string {
	rows := make([]string, numRows)
	direction := 1
	currentRow := 0
	for _, char := range s {
		rows[currentRow] += string(char)
		if currentRow == 0 {
			direction = 1
		}
		if currentRow == numRows-1 {
			direction = -1
		}
		currentRow += direction
	}
	return strings.Join(rows, "")
}
func convert2(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([]strings.Builder, numRows)
	direction := 1
	currentRow := 0
	for _, char := range s {
		rows[currentRow].WriteRune(char)
		if currentRow == 0 {
			direction = 1
		}
		if currentRow == numRows-1 {
			direction = -1
		}
		currentRow += direction
	}
	var result strings.Builder
	for _, group := range rows {
		result.WriteString(group.String())
	}
	return result.String()
}
