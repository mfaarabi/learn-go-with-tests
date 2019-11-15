package main

import "fmt"

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
	lengthOfNumbers := len(numbers)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbers {
		sums[i] = Sum(numbers)
	}

	return sums
}

func main() {
	fmt.Println(Hello("world"))
}
