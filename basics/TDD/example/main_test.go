package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	got := WaveSomeone("")
	expected := "Hello, World"
	if got != expected {
		t.Errorf("Got: %s but was expecting: %s", got, expected)
	}
}

func TestHello(t *testing.T) {
	got := WaveSomeone("Chris")
	expected := "Hello, Chris"

	if got != expected {
		t.Errorf("got %q expecting %q", got, expected)
	}
}
