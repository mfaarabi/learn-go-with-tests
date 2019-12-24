package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	width := rectangle.Width
	height := rectangle.Height
	return 2 * (width + height)
}

func (r Rectangle) Area() float64 {
	width := r.Width
	height := r.Height
	return width * height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	radius := c.Radius
	return math.Pi * math.Pow(radius, 2)
}
