package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	width := rectangle.Width
	height := rectangle.Height
	return 2 * (width + height)
}

func Area(rectangle Rectangle) float64 {
	width := rectangle.Width
	height := rectangle.Height
	return width * height
}
