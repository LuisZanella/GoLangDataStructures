package main

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	t.Run("Clone a text 6 times", func(t *testing.T) {
		res := CloneText("Hello", 6)
		expected := "HelloHelloHelloHelloHelloHelloHello"
		if res != expected {
			fmt.Errorf("Received: %s but expected was: %s", res, expected)
		}

	})
}

func BenchmarkRepeat(t *testing.B) {
	for i := 0; i < t.N; i++ {
		CloneText("a", i)
	}
	// fatal error: out of memory -> When is not good
	// when passes : 57087 109321 ns/op nano seconds per operation
}
