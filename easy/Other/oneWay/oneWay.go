package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compressStringChatGPT(text string) string {
	var compressed bool = true
	var compacted strings.Builder
	count := 1

	for i := 1; i < len(text); i++ {
		if text[i] == text[i-1] {
			count++
		} else {
			compacted.WriteString(string(text[i-1]))
			compacted.WriteString(strconv.Itoa(count))
			count = 1
		}
		if count > 1 {
			compressed = false
		}
	}

	compacted.WriteString(string(text[len(text)-1]))
	compacted.WriteString(strconv.Itoa(count))

	if compressed {
		return text
	}
	return compacted.String()
}

func myCompress(text string) string {
	var compressed strings.Builder
	counter := 0

	for i := 0; i < len(text); i++ {
		counter++
		if i+1 == len(text) || text[i] != text[i+1] {
			compressed.WriteString(string(text[i]))
			if counter > 1 {
				compressed.WriteString(strconv.Itoa(counter))
			}
			counter = 0
		}
	}
	compressedString := compressed.String()
	if len(compressedString) <= len(text) {
		return compressedString
	}
	return text
}

func main() {
	result := compressStringChatGPT("aaBBcccDDDbbBBQddd")
	fmt.Println("Result", result)

	result = compressStringChatGPT("abc")
	fmt.Println("Result", result)

	result = compressStringChatGPT("ab    c")
	fmt.Println("Result", result)

	result = compressStringChatGPT("abccd")
	fmt.Println("Result", result)

	result = myCompress("aaBBcccDDDbbBBQddd")
	fmt.Println("Result", result)

	result = myCompress("abc")
	fmt.Println("Result", result)

	result = myCompress("ab    c")
	fmt.Println("Result", result)

	result = myCompress("abccd")
	fmt.Println("Result", result)
}
