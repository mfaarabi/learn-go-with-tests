package main

import "testing"

func TestHello(t *testing.T) {
  expected := "Hello, Farabi"
  actual := Hello("Farabi")

  if expected != actual {
    t.Errorf("Expected %q, actual %q.", expected, actual)
  }
}