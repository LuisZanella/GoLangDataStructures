package main

import "testing"

func assert(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got: %+q Want: %+q", got, want)
	}
}
func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		name  string
		value string
		want  string
	}{
		{"Empty", "", ""},
		{"Not palindrome", "art", "a"},
		{"Pair palindrome", "babad", "bab"},
		{"Odd palindrome", "cbbd", "bb"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := longestPalindrome(c.value)
			assert(t, got, c.want)
		})

	}
}
