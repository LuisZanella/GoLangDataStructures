package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const (
	write = "write"
	sleep = "sleep"
)

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

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
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
