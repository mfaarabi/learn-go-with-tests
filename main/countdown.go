package main

import (
	"os"
	"time"

	"github.com/mfaarabi/learn-go-with-tests/countdown"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &DefaultSleeper{}
	countdown.Countdown(os.Stdout, sleeper)
}
