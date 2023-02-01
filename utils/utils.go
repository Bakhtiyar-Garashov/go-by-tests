package utils

import "testing"

func AssertCorrectMessage(t *testing.T, expected, got string) {
	t.Helper()

	if expected != got {
		t.Errorf("Expected %q, got %q", expected, got)
	}
}
