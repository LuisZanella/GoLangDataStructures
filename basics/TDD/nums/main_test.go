package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Positive sum", func(t *testing.T) {
		res := Sum(2, 5)
		expected := 7
		if res != expected {
			t.Errorf("Received: %d but %d was expected", res, expected)
		}
	})
}
