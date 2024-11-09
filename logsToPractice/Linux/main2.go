package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("Linux.log")
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
