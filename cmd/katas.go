// Print statistics about katas you've done.
package main

import (
	"flag"
	"log"

	"github.com/jreisinger/gokatas"
)

var showLastDoneDaysAgo = flag.Int("d", -1, "show only katas last done `days` ago or less")
var showLevel = flag.String("l", "", "show only katas of `level`")
var sortByColumn = flag.Int("c", 1, "sort by `column`")

func main() {
	flag.Parse()

	log.SetPrefix("gokatas: ")
	log.SetFlags(0)

	katas, err := gokatas.Get()
	if err != nil {
		log.Fatal(err)
	}
	gokatas.Print(katas, *showLastDoneDaysAgo, *sortByColumn, *showLevel)
}
