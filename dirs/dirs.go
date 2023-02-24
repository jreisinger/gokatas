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

	createEmptyFile(filepath.Join(tmp, "f1"))
	createEmptyFile(filepath.Join(tmp, "b", "f1"))
	createEmptyFile(filepath.Join(tmp, "b", "f2"))
	createEmptyFile(filepath.Join(tmp, "b", "c", "f3"))

	// Read directory entries.
	entries, err := os.ReadDir(tmp)
	if err != nil {
		log.Print(err)
	}
	for _, entry := range entries {
		fmt.Printf("--- %s ---\n", entry.Name())
		fi, err := entry.Info()
		if err != nil {
			log.Fatal(err)
		}
		printFileInfo(fi)
	}

	fmt.Println()

	// Walk directory recursively.
	err = filepath.WalkDir(tmp, visit)
	if err != nil {
		log.Printf("error walking the path %q: %v\n", tmp, err)
	}
}

func visit(path string, entry fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Printf("--- %s ---\n", path)
	fi, err := entry.Info()
	if err != nil {
		return err
	}
	printFileInfo(fi)
	return nil
}

var createEmptyFile = func(name string) {
	d := []byte("")
	err := os.WriteFile(name, d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func printFileInfo(fi fs.FileInfo) {
	fmt.Printf("IsDir: %v\nName: %v\nSize: %v\n",
		fi.IsDir(), fi.Name(), fi.Size())
}
