package older_test

import (
	"testing"
	"testing/fstest"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/jreisinger/gokatas/older"
)

func TestFilesReturnsFilesOlderThanGivenDuration(t *testing.T) {
	t.Parallel()
	now := time.Now()
	fsys := fstest.MapFS{
		"now":                       {ModTime: now},
		"subfolder/one-minute-old":  {ModTime: now.Add(-time.Minute)},
		"subfolder2/now":            {ModTime: now},
		"subfolder2/one-minute-old": {ModTime: now.Add(-time.Minute)},
	}
	want := []string{
		"subfolder/one-minute-old",
		"subfolder2/one-minute-old",
	}
	got := older.Files(fsys, 10*time.Second)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
