// Print statistics about katas you've done.
package main

import (
	"flag"
	"log"
	"time"

	"github.com/jreisinger/gokatas"
)

var sortByColumn = flag.Int("c", 1, "sort katas by `column`")
var lastDoneDaysAgo = flag.Int("d", daysSinceGoBirth(), "show only katas last done `days` ago or less")

func daysSinceGoBirth() int {
	birth := time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)
	days := time.Since(birth).Hours() / 24
	return int(days)
}

func main() {
	flag.Parse()

	log.SetPrefix("gokatas: ")
	log.SetFlags(0)

	katas, err := gokatas.Done(*lastDoneDaysAgo)
	if err != nil {
		log.Fatal(err)
	}
	gokatas.Print(katas, *sortByColumn)
}
