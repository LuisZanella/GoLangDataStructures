package main

import (
	"math"
	"testing"
)

func TestCalculatePerimeter(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		res := rectangle.CalculatePerimeter()
		expected := 40.0
		if res != expected {
			t.Errorf("Received: %.2f but %.2f was expected", res, expected)
		}
	})

}

func TestCalculateArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", &Rectangle{12, 6}, 72.0},
		{"Circle", &Circle{10}, 10.0 * 10.0 * math.Pi},
		{"Triangle", &Triangle{12, 6}, 36.0},
	}

	checkArea := func(t testing.TB, name string, shape Shape, expected float64) {
		t.Helper()
		got := shape.CalculateArea()
		if got != expected {
			t.Errorf("%v Received: %.2f but %.2f was expected", name, got, expected)
		}
	}

	t.Run("Area tests", func(t *testing.T) {
		for _, value := range areaTests {
			checkArea(t, value.name, value.shape, value.want)
		}
	})
}
