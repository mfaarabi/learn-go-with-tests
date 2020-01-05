package countdown

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountdown(t *testing.T) {
	t.Run("Prints countdown and final word", func(t *testing.T) {
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
	})

	t.Run("Sleep is called before print", func(t *testing.T) {
		spySleeperPrinter := &CountdownOperationSpy{}
		expected := []string{
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		Countdown(spySleeperPrinter, spySleeperPrinter)
		actual := spySleeperPrinter.Calls

		assert.Equal(t, expected, actual)
	})
}
