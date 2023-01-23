// Package countgo counts files with extension ".go" in fs.FS. It also shows
// how to use fs.FS in tests.
//
// When you’re dealing with trees of files, use fs.FS rather than trying to
// write the recursion code yourself.
//
// Based on bitfieldconsulting.com/golang/filesystems.
//
// Level: intermediate
// Topics: io/fs, testing/fstest, path/filepath
package countgo

import (
	"io/fs"
	"path/filepath"
)

func Files(fsys fs.FS) (count int) {
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		// NOTE: When there is an error walking the fsys it's stored in
		// err. We just ignore it. You might want to handle it in prod.
		if filepath.Ext(path) == ".go" {
			count++
		}
		return nil
	})
	return count
}
