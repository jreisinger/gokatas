package countgo

import (
	"os"
	"testing"
	"testing/fstest"
)

func TestFilesOnDisk(t *testing.T) {
	t.Parallel()
	fsys := os.DirFS("testdata")
	want := 4
	if got := Files(fsys); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFilesInMemory(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file.go":               {},
		"subfolder/file.go":     {},
		"subfolder2/another.go": {},
		"subfolder2/file.go":    {},
	}
	want := 4
	if got := Files(fsys); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
