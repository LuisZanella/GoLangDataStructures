package main

import (
	"bytes"
	"fmt"
)

const englishHelloPrefix = "Hello, "

func WaveSomeone(name string) string {
	var builder bytes.Buffer
	builder.WriteString(englishHelloPrefix)
	if len(name) == 0 {
		builder.WriteString("World")
	} else {
		builder.WriteString(name)
	}
	return builder.String()
}
func main() {
	fmt.Println(WaveSomeone(""))
}
