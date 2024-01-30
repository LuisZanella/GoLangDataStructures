package main

import (
	f "fmt"
	strings "strings"
)

func main() {
	var word string = "Luis"
	f.Println(ReverseWord(word))
	f.Println(ReverseWordSecondWay(word))
}

func ReverseWord(word string) (wordReversed string) {
	array := strings.Split(word, "")
	var backwards []string
	for i := len(array) - 1; i > -1; i-- {
		backwards = append(backwards, array[i])
	}
	wordReversed = strings.Join(backwards[:], "")
	return wordReversed
}
func ReverseWordSecondWay(word string) (wordReversed string) {
	var backwards []string
	for i := len(word); i > 0; i-- {
		backwards = append(backwards, string(word[i-1]))
	}
	wordReversed = strings.Join(backwards[:], "")
	return wordReversed
}
