package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
	}
	for _, test := range cases {
		t.Run(testName(test.time), func(t *testing.T) {
			got := SecondsInRadians(test.time)
			if got != test.angle {
				t.Errorf("Wanted: %v Got: %v ", test.angle, got)
			}
		})
	}
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
	want := math.Pi
	got := SecondsInRadians(thirtySeconds)
	if want != got {
		t.Fatalf("Wanted: %v Got: %v", want, got)
	}
}
func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
