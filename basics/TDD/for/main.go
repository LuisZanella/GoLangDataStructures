package main

import "bytes"

func CloneText(text string, times int) string {
	var builder bytes.Buffer
	for i := 0; i < times; i++ {
		builder.WriteString(text)
	}
	return builder.String()
}

func main() {

}
