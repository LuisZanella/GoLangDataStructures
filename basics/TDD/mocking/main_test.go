package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountDown(t *testing.T) {

	t.Run("Test without a spy", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		CountDown(buffer, &DefaultSleeper{})
		got := buffer.String()
		want := "3\n2\n1\nGo!"
		assertString(t, got, want)
	})
	t.Run("Test with spy", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		wantSpy := 3
		CountDown(buffer, spySleeper)
		got := buffer.String()
		want := "3\n2\n1\nGo!"

		assertNumber(t, spySleeper.Calls, wantSpy)
		assertString(t, got, want)
	})

}

func assertNumber(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("Got: %d Want: %d", got, want)
	}
}

func assertString(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("Got: %q Want: %q", got, want)
	}
}
