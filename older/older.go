// Package older finds files older than a given duration.
//
// Based on bitfieldconsulting.com/golang/filesystems.
//
// Level: intermediate
// Topics: io/fs, testing/fstest, time.Duration, tpg-tools
package older

import (
	"io/fs"
	"time"
)

func Files(fsys fs.FS, age time.Duration) (paths []string) {
	threshold := time.Now().Add(-age)
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		info, err := d.Info()
		if err != nil || info.IsDir() {
			return nil // skip
		}
		if info.ModTime().Before(threshold) {
			paths = append(paths, p)
		}
		return nil
	})
	return paths
}
