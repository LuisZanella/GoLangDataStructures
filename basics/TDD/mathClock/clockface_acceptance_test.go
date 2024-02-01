package clockFace_test

import (
	"bytes"
	clockFace "mathClock"

	"testing"
	"time"
)

func TestSecondHandAtMidNight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
	want := clockFace.Point{X: 150, Y: 150 - 90}
	b := bytes.Buffer{}
	got := clockFace.SecondHand(&b, tm)
	if got != want {
		t.Errorf("Got: %v Want: %v", got, want)
	}
}
func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := clockFace.Point{X: 150, Y: 150 + 90}
	b := bytes.Buffer{}
	got := clockFace.SecondHand(&b, tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
