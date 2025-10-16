package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height

}

const PI = 3.14159265359

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return PI * math.Pow(c.Radius, 2)
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Height float64
	Base   float64
}

func (t Triangle) Area() (are float64) {
	return (t.Base * t.Height) * 0.5
}

func Perimeter(rectangle Rectangle) (perimeter float64) {
	return 2 * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) (area float64) {
	return rectangle.Height * rectangle.Width
}
func main() {

}
