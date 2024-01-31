package main

import "testing"

func assert(t *testing.T, got, want bool) {
	if got != want {
		t.Errorf("Got: %t Want: %t", got, want)
	}
}

func TestInsUnique(t *testing.T) {
	var want bool

	cases := []struct {
		name  string
		value string
		want  bool
	}{
		{name: "Empty", value: "", want: true},
		{name: "Repeated", value: "AAACCCDD", want: false},
		{name: "Numbers", value: "1232134", want: false},
		{name: "Not extended ASCII", value: "AJDWIñ1234!@#$", want: false},
		{name: "Too much letters", value: "", want: true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			want = c.want
			got := isUnique(c.value)
			assert(t, got, want)
		})
	}

}
