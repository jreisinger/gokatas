package countgo_test

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/jreisinger/gokatas/countgo"
)

func TestCountgoFiles(t *testing.T) {
	tests := []struct {
		fsys fs.FS
		want int
	}{
		{fstest.MapFS{}, 0},
		{fstest.MapFS{"file.go": {}}, 1},
		{fstest.MapFS{"dir.go": &fstest.MapFile{Mode: fs.ModeDir}}, 0},
	}
	for i, test := range tests {
		got := countgo.Files(test.fsys)
		if got != test.want {
			t.Errorf("test %d: got %d, want %d", i, got, test.want)
		}
	}
}
