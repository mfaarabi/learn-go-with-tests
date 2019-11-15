package main

import (
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, expected, actual string) {
		t.Helper()
		if expected != actual {
			t.Errorf("Expected %q, actual %q.", expected, actual)
		}
	}

	t.Run("Hello with params", func(t *testing.T) {
		expected := "Hello, Farabi"
		actual := Hello("Farabi")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Hello with empty string", func(t *testing.T) {
		expected := "Hello, World"
		actual := Hello("")
		assertCorrectMessage(t, expected, actual)
	})
}

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, expected, actual int) {
		t.Helper()
		if expected != actual {
			t.Errorf("Expected %q, actual %q.", expected, actual)
		}
	}

	t.Run("Slices", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		expected := 6
		actual := Sum(numbers)
		assertCorrectMessage(t, expected, actual)
	})
}

func TestSumAll(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, expected, actual []int) {
		t.Helper()
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected %q, actual %q.", expected, actual)
		}
	}

	t.Run("Test", func(t *testing.T) {
		expected := []int{3, 9}
		actual := SumAll([]int{1, 2}, []int{0, 9})
		assertCorrectMessage(t, expected, actual)
	})
}
