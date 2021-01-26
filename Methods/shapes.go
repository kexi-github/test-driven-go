package Methods

import "math"

func (r *RectangleArea)Perimeter() float64 {

	return 2*(r.Width+r.Height)
}

func (r *RectangleArea)Area() float64 {

	return r.Width * r.Height
}

type RectangleArea struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (c *Circle)Perimeter() float64 {

	return 2*math.Pi*c.Radius
}

func (c *Circle)Area() float64 {

	return math.Pi*c.Radius*c.Radius
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Triangle struct {
	Base float64
	Height float64
}

func (t *Triangle)Perimeter() float64 {

	return 0
}

func (t *Triangle)Area() float64 {

	return (t.Base * t.Height) * 0.5
}
