package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	expected := 40.0
	actual := Perimeter(10.0, 10.0)

	assert.Equal(t, expected, actual)
}

func TestArea(t *testing.T) {
	expected := 72.0
	actual := Area(12.0, 6.0)

	assert.Equal(t, expected, actual)
}
