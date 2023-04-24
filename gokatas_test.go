package gokatas

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseKata(t *testing.T) {
	dir := "testdata"
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range entries {
		level, topics, err := parseKata(filepath.Join(dir, entry.Name()))
		if err != nil {
			t.Errorf("test failed: %v", err)
		}

		wantLevel := "beginner"
		if level != wantLevel {
			t.Errorf("got %s, want %s", level, wantLevel)
		}

		wantTopics := []string{"nothing", "really"}
		if !equal(topics, wantTopics) {
			t.Errorf("got %v, want %v", topics, wantTopics)
		}
	}
}

// equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
