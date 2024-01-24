package main

import "testing"

func TestAddBinary(t *testing.T) {
	var b1 string
	var b2 string
	t.Run("First Greater", func(t *testing.T) {
		b1 = "1010"
		b2 = "100"
		got := addBinary(b1, b2)
		expected := "1110"
		validate(t, got, expected)
	})
	t.Run("Last Greater", func(t *testing.T) {
		b1 = "1000"
		b2 = "10001"
		got := addBinary(b1, b2)
		expected := "11001"
		validate(t, got, expected)
	})
	t.Run("Empty", func(t *testing.T) {
		b1 = "1010"
		b2 = ""
		got := addBinary(b1, b2)
		expected := "1010"
		validate(t, got, expected)
	})
	t.Run("All ones", func(t *testing.T) {
		b1 = "1111"
		b2 = "1111"
		got := addBinary(b1, b2)
		expected := "11110"
		validate(t, got, expected)
	})
}
func validate(t *testing.T, got string, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("Got %q Expected: %q", got, expected)
	}
}
