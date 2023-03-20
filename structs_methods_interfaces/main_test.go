package main

import "testing"

// shape interface
// any struct that has Area method that returns float64 satisfies the Shape interface
type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		// The f is for our float64 and the .2 means print 2 decimal places
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		// Dependng the shape that is passed in from the shape, it will call the Area function of that shape
		got := shape.Area()

		if got != want {
			// g will print a more precise decimal number in the error message (fmt options).
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{5.0, 10.0}

		checkArea(t, rectangle, 50.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}

		checkArea(t, circle, 314.1592653589793)
	})
}

func TestAreaTable(t *testing.T) {
	// This creates an anonymous struct slice
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		// The fields (name, shape, want) are optionally added
		// We can also do {Rectangle{12, 6}, 72.0}
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		// use tt.name as the test description
		t.Run(tt.name, func(t *testing.T) {

			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}
