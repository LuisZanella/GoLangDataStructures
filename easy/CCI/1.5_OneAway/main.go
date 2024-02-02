package main

import (
	"bytes"
	"strconv"
)

func Commpression(s string) string {
	var compression bytes.Buffer
	var lastCharacter int32
	var counter int
	for _, v := range s {
		counter++
		if v != lastCharacter {
			compression.WriteRune(v)
			lastCharacter = v
			compression.WriteString(strconv.Itoa(counter))
			counter = 0
		}
	}
	return compression.String()
}
