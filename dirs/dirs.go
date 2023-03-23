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
	log.SetFlags(0)
	log.SetPrefix(os.Args[0] + ": ")

	const tmp = "a"

	// Make directory.
	err := os.Mkdir(tmp, 0750)
	if err != nil && !errors.Is(err, fs.ErrExist) {
		log.Fatal(err)
	}

	// Remove all, like rm -rf.
	defer os.RemoveAll(tmp)

	// Make all directories, like mkdir -p.
	err = os.MkdirAll(filepath.Join(tmp, "b", "c"), 0750)
	if err != nil && !errors.Is(err, fs.ErrExist) {
		log.Fatal(err)
	}

	createEmptyFile := func(name string) {
		err := os.WriteFile(name, []byte(""), 0640)
		if err != nil {
			log.Fatal(err)
		}
	}

	createEmptyFile(filepath.Join(tmp, "f1"))
	createEmptyFile(filepath.Join(tmp, "b", "f1"))
	createEmptyFile(filepath.Join(tmp, "b", "f2"))
	createEmptyFile(filepath.Join(tmp, "b", "c", "f3"))

	// Read directory entries.
	fmt.Println("--- os.ReadDir ---")
	entries, err := os.ReadDir(tmp)
	if err != nil {
		log.Print(err)
	}
	for _, entry := range entries {
		fi, err := entry.Info()
		if err != nil {
			log.Print(err)
		}
		printFileInfo(filepath.Join(tmp, entry.Name()), fi)
	}

	// Walk directory recursively.
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
	err = filepath.WalkDir(tmp, visit)
	if err != nil {
		log.Printf("error walking the path %q: %v\n", tmp, err)
	}
}

func printFileInfo(path string, fi fs.FileInfo) {
	fmt.Printf("Path\t%v\nName\t%v\nIsDir\t%v\nSize\t%v\n\n",
		path, fi.Name(), fi.IsDir(), fi.Size())
}
