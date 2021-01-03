package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, hasPerimeter: 40.0},
		{name: "Circle", shape: Circle{Radius: 5.0}, hasPerimeter: 31.41592653589793},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.shape.Perimeter()
			if actual != tt.hasPerimeter {
				t.Errorf("%#v: actual %g, expected %g", tt.name, actual, tt.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.shape.Area()
			if actual != tt.hasArea {
				t.Errorf("%#v: actual %g, expected %g", tt.shape, actual, tt.hasArea)
			}
		})
	}
}
