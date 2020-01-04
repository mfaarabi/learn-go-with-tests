package countdown

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	expected := `3
2
1
Go!`

	Countdown(buffer)
	actual := buffer.String()

	assert.Equal(t, expected, actual)
}
