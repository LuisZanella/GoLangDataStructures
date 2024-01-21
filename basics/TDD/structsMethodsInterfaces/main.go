package main

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}
type Triangle struct {
	width  float64
	height float64
}
type Shape interface {
	CalculateArea() float64
}

func (r *Rectangle) CalculatePerimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c *Circle) CalculateArea() float64 {
	return c.radius * c.radius * math.Pi
}
func (r *Rectangle) CalculateArea() float64 {
	return r.width * r.height
}
func (t *Triangle) CalculateArea() float64 {
	return t.height * t.width / 2
}
