package main

import "testing"

func assert(t *testing.T, got, want bool) {
	if got != want {
		t.Errorf("\n Got: %t \n Want: %t", got, want)
	}
}

func TestIsPalindromePermutation(t *testing.T) {
	cases := []struct {
		value string
		want  bool
	}{
		{"Tact coa", true},
		{"Tact coaT", false},
		{"", true},
		{"Hello", true},
		{"Tact coa---", true},
	}
	for _, c := range cases {
		t.Run(c.value, func(t *testing.T) {
			got := isPalindromePermutation(c.value)
			assert(t, got, c.want)
		})
	}
}
