package clockFace

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHandPoint(t time.Time) Point {
	angle := SecondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
func SecondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / (float64(t.Second())))
}

const secondHandLength = 90
const clockCentreX = 150
const clockCentreY = 150

func SecondHand(t time.Time) Point {
	p := SecondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	return p
}
