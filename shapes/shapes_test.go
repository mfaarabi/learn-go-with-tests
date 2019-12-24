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
	checkArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		assert.Equal(t, shape.Area, expected)
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		expected := 72.0
		checkArea(t, rectangle, expected)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793
		checkArea(t, circle, expected)
	})
}
