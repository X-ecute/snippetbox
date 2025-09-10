package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, expected, actual T) {
	t.Helper()

	if expected != actual {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func StringContains(t *testing.T, actual, expectedSubstring string) {
	t.Helper()
	if !strings.Contains(actual, expectedSubstring) {
		t.Errorf("got: %q; expected to contain: %q", actual, expectedSubstring)
	}
}
