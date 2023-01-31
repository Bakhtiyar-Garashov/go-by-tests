package main

import "testing"

func TestGreetWorld(t *testing.T) {
	expected := "Hello, World!"

	got := GreetWorld()

	if expected != got {
		t.Errorf("got %q want %q", got, expected)
	}
}
