package main

import "testing"

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
		expected := "Hello, "
		actual := Hello("")
		assertCorrectMessage(t, expected, actual)
	})
}
