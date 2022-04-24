// Print statistics about katas you've done.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jreisinger/gokatas"
)

var showAll = flag.Bool("a", false, "show all katas (default is those last done within two weeks)")
var sortByCount = flag.Bool("c", false, "sort katas by count (default is by last done)")

func main() {
	flag.Parse()
	katas, err := gokatas.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "gokatas: %v\n", err)
		os.Exit(1)
	}
	gokatas.Print(katas, *showAll, *sortByCount)
}
