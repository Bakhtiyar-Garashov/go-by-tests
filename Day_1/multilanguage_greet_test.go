package main

import "testing"

func assertCorrectMessage(t *testing.T, expected, got string) {
	t.Helper()

	if expected != got {
		t.Errorf("Expected %q, got %q", expected, got)
	}
}

func TestMultilanguageGreet(t *testing.T) {
	t.Run("Greet Road Runner in Estonian", func(t *testing.T) {
		expected := "Tere tulemast, Road Runner!"
		got := MultilanguageGreet("Road Runner", "EE")
		assertCorrectMessage(t, expected, got)
	})

	t.Run("Greet Road Runner in Azerbaijani", func(t *testing.T) {
		expected := "Salam, Road Runner!"
		got := MultilanguageGreet("Road Runner", "AZ")
		assertCorrectMessage(t, expected, got)
	})

	t.Run("Greet Road Runner in English", func(t *testing.T) {
		expected := "Hello, Road Runner!"
		got := MultilanguageGreet("Road Runner", "EN")
		assertCorrectMessage(t, expected, got)
	})
}
