package loops

import "testing"

func TestRepeatChar(t *testing.T) {
	expected := "aaaaaa"

	got := RepeatChar("a", 6)

	if expected != got {
		t.Errorf("Expected %q, got %q", expected, got)
	}
}
