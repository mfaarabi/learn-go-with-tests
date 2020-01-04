package main

import (
	"fmt"
	"os"

	"github.com/mfaarabi/learn-go-with-tests/greet"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i < 5; i++ {
	for _, number := range numbers {
		// 	sum += numbers[i]
		sum += number
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int

	for _, numbers := range numbers {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func main() {
	fmt.Println(Hello("world"))
	greet.Greet(os.Stdout, "Wendy\n")
}
