package main

import "testing"

func TestIsSymetric(t *testing.T) {
	t.Run("", func(t *testing.T) {

	})
}

func assert(t *testing.T, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %t Want: %t", got, want)
	}
}
