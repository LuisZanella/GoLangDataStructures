package main

import (
	"fmt"
	"testing"
)

type testNumberToRoman struct {
	number   int
	expected string
}

func TestNumberToRoman(t *testing.T) {
	t.Run("0 value", func(t *testing.T) {
		got := IntToRoman(0)
		assertTestNumber(t, got, "")
	})
	t.Run("Common number", func(t *testing.T) {
		got := IntToRoman(3749)
		assertTestNumber(t, got, "MMMDCCXLIX")
	})
	subtractionCases := []testNumberToRoman{{number: 4, expected: "IV"}, {number: 9, expected: "IX"},
		{number: 40, expected: "XL"}, {number: 90, expected: "XC"}, {number: 400, expected: "CD"},
		{number: 900, expected: "CM"}}

	for _, subtractCase := range subtractionCases {
		t.Run(fmt.Sprintf("%d to %q subtraction", subtractCase.number, subtractCase.expected), func(t *testing.T) {
			got := IntToRoman(subtractCase.number)
			assertTestNumber(t, got, subtractCase.expected)
		})
	}
}

func assertTestNumber(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
