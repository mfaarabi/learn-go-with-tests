package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	expected := 40.0
	actual := Perimeter(Rectangle{10.0, 10.0})

	assert.Equal(t, expected, actual)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		actual := tt.shape.Area()
		assert.Equal(t, tt.expected, actual)
	}
}
