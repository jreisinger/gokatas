// Print statistics about katas you've done.
package main

import (
	"flag"
	"log"
	"os"
)

var sortByColumn = flag.Int("c", 1, "sort katas by `column`")
var gokatasRepo = flag.String("r", ".", "path to gokatas repository")

func main() {
	flag.Parse()

	log.SetPrefix("gokatas: ")
	log.SetFlags(0)

	if *gokatasRepo != "." {
		if err := os.Chdir(*gokatasRepo); err != nil {
			log.Fatal(err)
		}
	}

	katas, err := Get()
	if err != nil {
		log.Fatalf("getting katas: %v", err)
	}
	Print(katas, *sortByColumn)
}
