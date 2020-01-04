package countdown

import (
	"fmt"
	"io"
	"time"
)

const seconds = 3
const secondDelay = 1 * time.Second
const finalWord = "Go!"

func Countdown(out io.Writer) {
	for i := seconds; i > 0; i-- {
		time.Sleep(secondDelay)
		fmt.Fprintln(out, i)
	}
	time.Sleep(secondDelay)
	fmt.Fprint(out, finalWord)
}
