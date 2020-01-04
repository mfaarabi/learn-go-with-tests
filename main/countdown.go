package main

import (
	"os"

	"github.com/mfaarabi/learn-go-with-tests/countdown"
)

func main() {
	countdown.Countdown(os.Stdout)
}
