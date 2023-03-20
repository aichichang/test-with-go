package main

import "math"

// struct
type Rectangle struct {
	Width  float64
	Height float64
}

// method
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// circle struct
type Circle struct {
	Radius float64
}

// circle method
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

// triangle method
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
