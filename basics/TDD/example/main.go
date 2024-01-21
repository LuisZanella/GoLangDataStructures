package main

import (
	"bytes"
	"fmt"
)

var helloMap = map[string]string{"Spanish": "Hola", "French": "Bonjour", "English": "Hello"}

func helloLanguage(language string) string {
	if val, ok := helloMap[language]; ok {
		return val
	} else {
		return helloMap["English"]
	}
}

func WaveSomeone(name string, language string) string {
	var builder bytes.Buffer
	if name == "" {
		name = "World"
	}
	builder.WriteString(helloLanguage(language))
	builder.WriteString(", ")
	builder.WriteString(name)
	return builder.String()
}
func main() {
	fmt.Println(WaveSomeone("", ""))
}
