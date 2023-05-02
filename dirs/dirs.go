// Dirs shows how to work with filesystem directories.
// Based on https://gobyexample.com/directories.
//
// Level: beginner
// Topics: io/fs, directories, scripting
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// CLI tool style logging.
	log.SetFlags(0)
	log.SetPrefix(os.Args[0] + ": ")

	const a = "a"

	// Make directory.
	err := os.Mkdir(a, 0750)
	if err != nil && !errors.Is(err, os.ErrExist) {
		log.Fatal(err)
	}

	// Remove all (after we are done), like rm -rf.
	defer os.RemoveAll(a)

	// Make all directories, like mkdir -p.
	path := filepath.Join(a, "b", "c")
	if err := os.MkdirAll(path, 0750); err != nil {
		log.Fatal(err)
	}

	createEmptyFile := func(name string) {
		err := os.WriteFile(name, []byte(""), 0640)
		if err != nil {
			log.Fatal(err)
		}
	}

	createEmptyFile(filepath.Join(a, "1"))
	createEmptyFile(filepath.Join(a, "b", "1"))
	createEmptyFile(filepath.Join(a, "b", "2"))
	createEmptyFile(filepath.Join(a, "b", "c", "3"))

	// Read directory entries and print file info.
	fmt.Println("--- os.ReadDir ---")
	entries, err := os.ReadDir(a)
	if err != nil {
		log.Print(err)
	}
	for _, entry := range entries {
		fi, err := entry.Info()
		if err != nil {
			log.Print(err)
			continue
		}
		path := filepath.Join(a, entry.Name())
		printFileInfo(path, fi)
	}

	// Walk directory recursively and print file info.
	fmt.Println("--- filepath.WalkDir ---")
	visit := func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fi, err := entry.Info()
		if err != nil {
			return err
		}
		printFileInfo(path, fi)
		return nil
	}
	if err := filepath.WalkDir(a, visit); err != nil {
		log.Printf("walking %s: %v", a, err)
	}
}

func printFileInfo(path string, fi fs.FileInfo) {
	fmt.Printf("Path\t%v\nName\t%v\nIsDir\t%v\nSize\t%v\n\n",
		path, fi.Name(), fi.IsDir(), fi.Size())
}
