package main

import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	CountDown(buffer)
	got := buffer.String()
	want := "3"
	if got != want {
		t.Errorf("Got: %q Want: %q", got, want)
	}
}
