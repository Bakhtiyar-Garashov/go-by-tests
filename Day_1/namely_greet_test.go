package main

import "testing"

func TestNamelyGreet(t *testing.T) {
	t.Run("Function returns Hello World if the param is empty string", func(t *testing.T) {
		expected := "Hello, World!"
		got := NamelyGreet("")

		if expected != got {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	})

	t.Run("Function greets the name", func(t *testing.T) {
		expected := "Hello, Road Runner!"
		got := NamelyGreet("Road Runner")

		if expected != got {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	})
}
