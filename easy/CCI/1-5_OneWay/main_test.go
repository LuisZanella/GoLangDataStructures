package main

import "testing"

func assert(t *testing.T, got, want bool) {
	if got != want {
		t.Errorf("Got: %t Want: %t", got, want)
	}
}
func TestOneWay(t *testing.T) {
	cases := []struct {
		name   string
		values []string
		want   bool
	}{
		{"Empty", []string{"", ""}, true},
		{"Same", []string{"T E S T", "T E S T"}, true},
		{"Update", []string{"Hello", "hello"}, true},
		{"Insert", []string{"This is", "This isA"}, true},
		{"Remove", []string{"This is", "This i"}, true},
		{"Case sensitive", []string{"Helloo", "hello"}, false},
		{"Case sensitive Doble Update", []string{"Hello", "HeLlO"}, false},
		{"Case sensitive Remove", []string{"Hell", "hello"}, false},
		{"Totally Different", []string{"34234!dsaEW", "ADS@!cas"}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Oneway(c.values)
			assert(t, got, c.want)
		})

	}
}
