package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 10000000; i++ {
	}

	actualTime := time.Now()
	elapsed := actualTime.Sub(start)

	fmt.Printf("t : %s \n", elapsed)
}
