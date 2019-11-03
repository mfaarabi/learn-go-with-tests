package main

import "testing"

func TestHello(t *testing.T) {
  expected := "Hello, world"
  actual := Hello()

  if expected != actual {
    t.Errorf("Expected %q, actual %q.", expected, actual)
  }
}