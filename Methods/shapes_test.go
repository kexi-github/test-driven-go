package Methods_test

import (
	"math"
	"testing"
	"tests/Methods"
)

func TestPerimeter(t *testing.T){

	rect := Methods.RectangleArea{10.0,10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want{
		t.Errorf("got %.2f want %.2f",got,want)
	}
}

func TestArea(t *testing.T)  {

	areaTests := []struct{
		name string
		shape Methods.Shape
		want float64
	}{
		{"RectangleArea",&Methods.RectangleArea{12,6},72.0},
		{"Circle",&Methods.Circle{10},math.Pi*10*10},
		{"Triangle",&Methods.Triangle{12,6},36.00},
	}

	for _,tt := range areaTests{

		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want{
				t.Errorf("%v got %.2f want %.2f",tt.name,got,tt.want)
			}
		})
	}
}
