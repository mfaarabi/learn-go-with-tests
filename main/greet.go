package main

import (
	"os"

	"github.com/mfaarabi/learn-go-with-tests/greet"
)

func main() {
	greet.Greet(os.Stdout, "Wendy\n")
}
