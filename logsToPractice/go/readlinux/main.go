package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("../../../logsToPractice/Linux/Linux.log")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	printableChars := regexp.MustCompile(`[[:print:]]+`)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Error") || strings.Contains(line, "error") || strings.Contains(line, "Failure") {
			fields := strings.Fields(line)
			rawMessage := strings.Join(fields[3:], " ")
			sanitizedMessage := printableChars.FindAllString(rawMessage, -1)
			if len(fields) > 2 {
				fmt.Printf("Date: %s %s - Message: %s\n\n", fields[0], fields[1], strings.Join(sanitizedMessage, " "))
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file", err)
	}
}
