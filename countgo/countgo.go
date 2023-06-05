// Package countgo counts files with extension ".go" in fs.FS. It also shows how
// to use fs.FS in tests.
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
	visit := func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return fs.SkipDir
		}
		if !entry.IsDir() && filepath.Ext(path) == ".go" {
			count++
		}
		return nil
	}
	fs.WalkDir(fsys, ".", visit)
	return
}
