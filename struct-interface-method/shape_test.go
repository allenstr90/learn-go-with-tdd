package structinterfacemethod

import "testing"

func TestArea(t *testing.T) {
	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 3, Height: 2}, want: 6},
		{name: "Circle", shape: Circle{Radius: 3.5}, want: 38.48451000647496},
		{name: "Triangle", shape: Triangle{Base: 10, Height: 4}, want: 20},
	}
	for _, shapeArea := range areaTest {
		t.Run(shapeArea.name, func(t *testing.T) {
			got := shapeArea.shape.Area()

			if got != shapeArea.want {
				t.Errorf("%#v got %g but want %g", shapeArea.shape, got, shapeArea.want)
			}
		})
	}
}
