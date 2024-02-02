package main

import "testing"

func TestCompression(t *testing.T) {
	cases := []struct {
		name  string
		value string
		want  string
	}{
		{"1 size", "a", "a1"},
		{"empty", "", ""},
		{"single letter repeated", "aaaaaaaaaa", "a10"},
		{"multiple repeated", "aaaaaaaaaabbbbbbbcccaaa", "a10b7c3a3"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Commpression(c.value)
			assert(t, got, c.want)
		})
	}
}

func assert(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("\n Got:%q Want:%q", got, want)
	}
}
