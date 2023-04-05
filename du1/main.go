// The du1 command computes the disk usage of the files in a directory.
// Taken from https://github.com/adonovan/gopl.io/blob/master/ch8/du1/main.go.
//
// Level: intermediate
// Topics: concurrency, recursion, filesystem
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	// Determine the initial directories.
	flag.Parse()
	dirs := flag.Args()
	if len(dirs) == 0 {
		dirs = []string{"."}
	}

	// Walk the initial directories.
	fileSizes := make(chan int64)
	go func() {
		for _, dir := range dirs {
			walkDir(dir, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the number of files and bytes.
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fi, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "du1: %v\n", err)
			}
			fileSizes <- fi.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []fs.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
	}
	return entries
}
