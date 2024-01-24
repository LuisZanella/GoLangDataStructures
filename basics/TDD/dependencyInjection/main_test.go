package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("Adding testing", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")
		got := buffer.String()
		want := "Hello, Chris"
		if got != want {
			t.Errorf("Got: %q Want: %q", got, want)
		}
	})
}
