// Dirs shows how to work with filesystem directories.
// Based on https://gobyexample.com/directories.
//
// Level: beginner
// Topics: io/fs, directories, scripting
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix(os.Args[0] + ": ")

	// Make directory.
	err := os.Mkdir("a", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	// Remove all, like rm -rf.
	defer os.RemoveAll("a")

	// Make all directories, like mkdir -p.
	err = os.MkdirAll(filepath.Join("a", "b", "c"), 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	createEmptyFile(filepath.Join("a", "f1"))
	createEmptyFile(filepath.Join("a", "b", "f1"))
	createEmptyFile(filepath.Join("a", "b", "f2"))
	createEmptyFile(filepath.Join("a", "b", "c", "f3"))

	// Read directory entries.
	entries, err := os.ReadDir("a")
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
	err = filepath.WalkDir("a", visit)
	if err != nil {
		log.Printf("error walking the path %q: %v\n", "a", err)
	}
}

var createEmptyFile = func(name string) {
	d := []byte("")
	err := os.WriteFile(name, d, 0644)
	if err != nil {
		log.Fatal(err)
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

func printFileInfo(fi fs.FileInfo) {
	fmt.Printf("IsDir: %v\nName: %v\nSize: %v\n",
		fi.IsDir(), fi.Name(), fi.Size())
}
