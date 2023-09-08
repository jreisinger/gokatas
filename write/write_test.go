package write_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jreisinger/gokatas/write"
)

func TestWriteToFile_WritesToFile(t *testing.T) {
	t.Parallel()

	// path := "testdata/file.txt"
	// _, err := os.Stat(path)
	// if err == nil {
	// 	t.Fatalf("test artifact not cleaned up: %q", path)
	// }
	// defer os.Remove(path)
	path := t.TempDir() + "/write_test.txt"

	want := []byte("hello")
	if err := write.ToFile(path, want); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestWriteToFile_ReturnsErrorForUnwritableFile(t *testing.T) {
	t.Parallel()
	path := "bogusdir/file.txt"
	err := write.ToFile(path, []byte{})
	if err == nil {
		t.Fatal("want error when file not writable")
	}
}

// NOTE: tests also serve to *document* the required behavior. A wise
// programmer will read all the tests covering the unit of interest before
// attempting to refactor it. This refactoring might happen years from now ...
// Gandalf spent 17 years researching the ring.
func TestWriteToFile_ClobbersExistingFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/clobber_test.txt"
	err := os.WriteFile(path, []byte{1, 2, 3}, 0o600)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriteToFile_SetsCorrectFilePermissions(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/write_test.txt"
	if err := write.ToFile(path, []byte{}); err != nil {
		t.Fatal(err)
	}
	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	perm := stat.Mode().Perm()
	if perm != 0o600 {
		t.Errorf("want file mode 0o600, got 0%o", perm)
	}
}

// To prevent a pre-population attack.
func TestWriteToFile_SetsCorrectFilePermissionsOnExistingFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/perms_test.txt"
	// Pre-create empty file with open perms
	if err := os.WriteFile(path, []byte{}, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := write.ToFile(path, []byte{1, 2, 3}); err != nil {
		t.Fatal(err)
	}
	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	perm := stat.Mode().Perm()
	if perm != 0o600 {
		t.Errorf("want 0o600, got 0o%o", perm)
	}
}
