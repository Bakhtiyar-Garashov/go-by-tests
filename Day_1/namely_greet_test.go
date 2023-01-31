package main

import "testing"

func TestNamelyGreet(t *testing.T) {
	expected := "Hello, Road Runner!"
	got := NamelyGreet("Road Runner")

	if expected != got {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}
