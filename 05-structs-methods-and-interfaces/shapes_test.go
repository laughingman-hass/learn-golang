package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 10.0, Height: 10.0}, want: 40.0},
        {shape: Circle{Radius: 5.0}, want: 31.41592653589793},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g, want %g", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12.0, Height: 6.0}, want: 72.0},
		{shape: Circle{Radius: 10.0}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g, want %g", got, tt.want)
		}
	}
}
