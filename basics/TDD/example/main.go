package hello

import "fmt"

func HelloWorld() string {
	return "Hello, World"
}
func WaveSomeone(name string) string {
	return "Hello, " + name
}
func main() {
	fmt.Println(HelloWorld())
}
