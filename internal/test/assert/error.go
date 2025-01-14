package assert

import (
	"strings"
	"testing"
)

// NoError checks if there was no error provided.
func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

// NoError2 checks if there was no error provided and returns a value.
func NoError2[T any](t *testing.T) func(T, error) T {
	t.Helper()
	return func(r T, err error) T {
		t.Helper()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		return r
	}
}

// Contains checks if the specified substring is found in data.
func Contains[T ~string, K ~string](t *testing.T, data T, substring K) {
	t.Helper()
	if !strings.Contains(
		string(data),
		string(substring),
	) {
		t.Fatalf("Expected substring '%s' not found in '%s'", substring, data)
	}
}
