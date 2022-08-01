// Package countgo counts files with extension ".go" in fs.FS. Based on
// bitfieldconsulting.com/golang/filesystems.
package countgo

import (
	"io/fs"
	"path/filepath"
)

func Files(fsys fs.FS) (count int) {
	fn := func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".go" {
			count++
		}
		return nil
	}
	fs.WalkDir(fsys, ".", fn)
	return count
}
