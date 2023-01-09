package countgo_test

import (
	"os"
	"testing"
	"testing/fstest"

	"github.com/jreisinger/gokatas/countgo"
)

func TestFilesOnDisk(t *testing.T) {
	t.Parallel()
	fsys := os.DirFS("testdata")
	want := 4
	if got := countgo.Files(fsys); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFilesInMemory(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"subfolder/file.go":     {},
		"subfolder2/another.go": {},
		"subfolder2/file.go":    {},
		"file.go":               {},
	}
	want := 4
	if got := countgo.Files(fsys); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
