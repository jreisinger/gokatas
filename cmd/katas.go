// Print statistics about katas you've done.
package main

import (
	"flag"
	"log"

	"github.com/jreisinger/gokatas"
)

var showLastDoneDaysAgo = flag.Int("d", -1, "show only katas last done `days` ago or less")
var sortByCount = flag.Bool("c", false, "sort katas by done count")
var level = flag.String("l", "", "print only katas of `level`")

func main() {
	flag.Parse()

	log.SetPrefix("gokatas: ")
	log.SetFlags(0)

	katas, err := gokatas.Get()
	if err != nil {
		log.Fatal(err)
	}
	gokatas.Print(katas, *showLastDoneDaysAgo, *sortByCount, *level)
}
