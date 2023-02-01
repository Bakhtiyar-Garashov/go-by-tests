package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	t.Run("Addition", func(t *testing.T) {
		expected := 4.5

		result := Calculate(1.5, 3.0, "+")

		if result != expected {
			t.Errorf("expected '%f' but got '%f'", expected, result)
		}

	})

	t.Run("Subtraction", func(t *testing.T) {
		expected := 2.0

		result := Calculate(3.0, 1.0, "-")

		if result != expected {
			t.Errorf("expected '%f' but got '%f'", expected, result)
		}

	})

	t.Run("Division", func(t *testing.T) {
		expected := 1.5

		result := Calculate(3.0, 2.0, "/")

		if result != expected {
			t.Errorf("expected '%f' but got '%f'", expected, result)
		}

	})

	t.Run("Multiplication", func(t *testing.T) {
		expected := 4.5

		result := Calculate(3.0, 1.5, "*")

		if result != expected {
			t.Errorf("expected '%f' but got '%f'", expected, result)
		}

	})

	t.Run("Default case", func(t *testing.T) {
		expected := 0.0

		result := Calculate(3.0, 1.5, "%")

		if result != expected {
			t.Errorf("expected '%f' but got '%f'", expected, result)
		}

	})
}
