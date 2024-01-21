package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		received := WaveSomeone("Chris", "")
		expected := "Hello, Chris"

		assertCorrectMessage(t, received, expected)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		received := WaveSomeone("", "")
		expected := "Hello, World"

		assertCorrectMessage(t, received, expected)
	})
	t.Run("in Spanish", func(t *testing.T) {
		received := WaveSomeone("Elodie", "Spanish")
		expected := "Hola, Elodie"
		assertCorrectMessage(t, received, expected)
	})
}

func assertCorrectMessage(t testing.TB, received, expected string) {
	t.Helper()
	if received != expected {
		t.Errorf("received %q expected %q", received, expected)
	}
}
