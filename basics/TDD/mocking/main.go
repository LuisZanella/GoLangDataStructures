package main

import (
	"fmt"
	"io"
	"os"
)

func CountDown(out io.Writer) {
	fmt.Fprint(out, "3")
}

func main() {
	CountDown(os.Stdout)
}
