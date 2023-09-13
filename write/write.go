// Package write shows how to write programs manipulating files with test-first
// approach (TDD). Adapted from: https://bitfieldconsulting.com/books/tools
//
// Level: Intermediate
// Topics: TDD, files, go-cmp, tpg-tools
package write

import (
	"io/fs"
	"os"
)

func ToFile(path string, data []byte) error {
	var mode fs.FileMode = 0o600
	if err := os.WriteFile(path, data, mode); err != nil {
		return err
	}
	// NOTE: It’s still possible that the attacker might be able to read at
	// least some data between the time WriteFile starts executing and the
	// time Chmod closes the permissions, though. So for secret data that
	// really matters, we’d need to implement writing to files at a lower
	// level, and that’s beyond the scope of this program.
	return os.Chmod(path, mode)
}
