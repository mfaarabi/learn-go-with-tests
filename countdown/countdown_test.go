package countdown

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	expected := `3
2
1
Go!`

	Countdown(buffer, spySleeper)
	actual := buffer.String()

	assert.Equal(t, expected, actual)
	assert.Equal(t, 4, spySleeper.Calls)
}
