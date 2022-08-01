// Package countgo counts ".go" files in fs.FS. Based on
// bitfieldconsulting.com/golang/filesystems.
package countgo

import (
	"io/fs"
	"path/filepath"
)

func Files(fsys fs.FS) (count int) {
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ".go" {
			count++
		}
		return nil
	})
	return count
}
