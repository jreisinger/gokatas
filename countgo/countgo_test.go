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
	got := Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
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
	got := Files(fsys)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
