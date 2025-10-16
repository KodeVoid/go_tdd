package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)

	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}

func TestArea(t *testing.T) {
	checkArea := func(shape Shape, want float64, t testing.TB) {
		t.Helper()
		got := shape.Area()
		epsilon := 1e-9

		if math.Abs(got-want) > epsilon {
			t.Errorf("got %.2f  want %.2f", got, want)
		}

	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 100.0
		checkArea(rectangle, want, t)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}

		want := 314.1592653589793

		checkArea(circle, want, t)
	})

	t.Run("table tests", func(t *testing.T) {
		areaTests := []struct {
			name  string
			shape Shape
			want  float64
		}{
			{"rectangle 12x6", Rectangle{12, 6}, 72.0},
			{"square 10x10", Rectangle{10, 10}, 100.0},
			{"circle r10", Circle{10}, 314.1592653589793},
			{"Triangle", Triangle{12, 6}, 36.0},
		}

		for _, tt := range areaTests {
			t.Run(tt.name, func(t *testing.T) {
				checkArea(tt.shape, tt.want, t)
			})
		}
	})

}
