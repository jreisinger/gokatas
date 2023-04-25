package countgo

import (
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestFiles(t *testing.T) {
	tests := []struct {
		fsys fs.FS
		want int
	}{
		{fstest.MapFS{}, 0},
		{fstest.MapFS{"dir.go": &fstest.MapFile{Mode: fs.ModeDir}}, 0},
		{fstest.MapFS{"dir/file.go": {}, "file": {}, "file.go": {}}, 2},
	}

	for i, test := range tests {
		got := Files(test.fsys)
		if got != test.want {
			t.Errorf("test %d: got %d, want %d", i, got, test.want)
		}
	}
}
