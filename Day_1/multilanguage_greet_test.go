package main

import (
	"testing"

	utils "github.com/Bakhtiyar-Garashov/go-by-tests/utils"
)

func TestMultilanguageGreet(t *testing.T) {
	t.Run("Greet Road Runner in Estonian", func(t *testing.T) {
		expected := "Tere tulemast, Road Runner!"
		got := MultilanguageGreet("Road Runner", "EE")
		utils.AssertCorrectMessage(t, expected, got)
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
