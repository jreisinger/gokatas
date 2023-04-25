// Print statistics about katas you've done.
package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/jreisinger/gokatas"
)

var sortByColumn = flag.Int("c", 1, "sort katas by `column`")
var lastDoneDaysAgo = flag.Int("d", daysSinceGoBirth(), "show only katas last done `days` ago or less")
var gokatasRepo = flag.String("r", ".", "path to gokatas repository")

func daysSinceGoBirth() int {
	birth := time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)
	days := time.Since(birth).Hours() / 24
	return int(days)
}

func main() {
	flag.Parse()

	log.SetPrefix("gokatas: ")
	log.SetFlags(0)

	if *gokatasRepo != "." {
		if err := os.Chdir(*gokatasRepo); err != nil {
			log.Fatal(err)
		}
	}

	katas, err := gokatas.Get(*lastDoneDaysAgo)
	if err != nil {
		log.Fatalf("getting katas: %v", err)
	}
	gokatas.Print(katas, *sortByColumn)
}
