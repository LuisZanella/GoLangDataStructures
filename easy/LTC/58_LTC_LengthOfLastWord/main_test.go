package main

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	var word string
	t.Run("Two Words", func(t *testing.T) {
		word = "Hello World"
		got := lengthOfLastWord(word)
		expect := 5
		if got != expect {
			t.Errorf("Recevived %d but expected was: %d", got, expect)
		}
	})
	t.Run("One Word", func(t *testing.T) {
		word = "World"
		got := lengthOfLastWord(word)
		expect := 5
		if got != expect {
			t.Errorf("Recevived %d but expected was: %d", got, expect)
		}
	})
	t.Run("Three words with extra spaces", func(t *testing.T) {
		word = "         Hello      World                "
		got := lengthOfLastWord(word)
		expect := 5
		if got != expect {
			t.Errorf("Recevived %d but expected was: %d", got, expect)
		}
	})
	t.Run("One Word extra spaces", func(t *testing.T) {
		word = "         World"
		got := lengthOfLastWord(word)
		expect := 5
		if got != expect {
			t.Errorf("Recevived %d but expected was: %d", got, expect)
		}
	})
}
