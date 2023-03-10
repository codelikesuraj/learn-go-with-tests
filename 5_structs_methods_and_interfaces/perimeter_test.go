package main

import "math"
import "testing"

// func TestPerimeter (t *testing.T) {
// 	perimeterTests := []struct {
// 		name string
// 		shape Shape
// 		want float64
// 	}{
// 		{name: "Circle", shape: Circle{20.00}, want: 2 * math.Pi * 20.00},
// 		{name: "Rectangle", shape: Rectangle{30.00, 20.0}, want: 100.00},
// 	}

// 	for _, tt := range perimeterTests {
// 		t.Run(tt.name, func(t *testing.T) {	
// 			got := tt.shape.Perimeter()
// 			if got != tt.want {
// 				t.Errorf("got %g want %g", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestArea (t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		want float64
	}{
		{name: "Circle", shape: Circle{10}, want: math.Pi * 10 * 10},
		{name: "Rectangle", shape: Rectangle{10.00, 30.00}, want: 300.00},
		{name: "Triangle", shape: Triangle{10.00, 30.00}, want: 150},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}