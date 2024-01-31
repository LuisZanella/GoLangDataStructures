package main

import "testing"

func assert(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("\n Got: %q  \n Want: %q", got, want)
	}
}

func TestURLify(t *testing.T) {
	cases := []struct {
		name        string
		value       string
		valueLength int
		want        string
	}{
		{"Word with spaces", "Mr John Smith    ", 13, "Mr%20John%20Smith"},
		{"Empty", "", 0, ""},
		{"One space", "   ", 1, "%20"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := urlLify(c.value, c.valueLength)
			assert(t, got, c.want)
		})
	}
}
