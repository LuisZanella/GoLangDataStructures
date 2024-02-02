package main

import (
	"bytes"
	"strconv"
)

func Commpression(s string) string {
	var compression bytes.Buffer
	var lastCharacter int32
	var counter int
	for i, v := range s {
		if lastCharacter == 0 {
			lastCharacter = v
		}
		if i == len(s)-1 {
			counter++
		}
		if v != lastCharacter || i == len(s)-1 {
			compression.WriteRune(lastCharacter)
			compression.WriteString(strconv.Itoa(counter))
			lastCharacter = v
			counter = 0
		}
		counter++
	}
	return compression.String()
}
