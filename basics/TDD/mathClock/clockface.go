package clockFace

import (
	"fmt"
	"io"
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength = 90
	minuteHandLength = 80
	clockCentreX     = 150
	clockCentreY     = 150
	svgStart         = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
	bezel  = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
	svgEnd = `</svg>`
)

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / (float64(t.Second())))
}
func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func handTag(w io.Writer, p Point) {
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	SecondHand(w, t)
	minuteHand(w, t)
	io.WriteString(w, svgEnd)
}
func minutesInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Minute()))
}
func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}
func SecondHand(w io.Writer, t time.Time) Point {
	p := makeHand(secondHandPoint(t), secondHandLength)
	handTag(w, p)
	return p
}

func minuteHand(w io.Writer, t time.Time) Point {
	p := makeHand(minuteHandPoint(t), minuteHandLength)
	handTag(w, p)
	return p
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	return Point{p.X + clockCentreX, p.Y + clockCentreY}
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / 12) +
		(math.Pi / (6 / float64(t.Hour()%12)))
}
func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}
