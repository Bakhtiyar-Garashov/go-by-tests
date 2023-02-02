package main

import "testing"

func TestSumArray(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}

		expected := 15

		got := SumArray(input)

		if expected != got {
			t.Errorf("Expected %d, got %d", expected, got)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumArray(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
