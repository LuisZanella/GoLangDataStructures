package main

import (
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	t.Run("Empty String is received", func(t *testing.T) {
		got := FindSubstring("", []string{""})
		assertArrayHelper(t, got, []int{})
	})
	t.Run("Two Substrings", func(t *testing.T) {
		got := FindSubstring("barfoothefoobarman", []string{"foo", "bar"})
		assertArrayHelper(t, got, []int{0, 9})
	})
	t.Run("Extra letters", func(t *testing.T) {
		got := FindSubstring("asdbarfoothefoobarman", []string{"foo", "bar"})
		assertArrayHelper(t, got, []int{3, 12})
	})
	t.Run("Invalid indexes", func(t *testing.T) {
		got := FindSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"})
		assertArrayHelper(t, got, []int{})
	})

}

func assertArrayHelper(t *testing.T, got, want []int) {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
