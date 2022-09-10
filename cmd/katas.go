// Print statistics about katas you've done.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jreisinger/gokatas"
)

var showLastDoneDaysAgo = flag.Int("d", 14, "show only katas last done `days` ago or less")
var sortByCount = flag.Bool("c", false, "sort katas by done count")

func main() {
	flag.Parse()
	katas, err := gokatas.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "gokatas: %v\n", err)
		os.Exit(1)
	}
	gokatas.Print(katas, *showLastDoneDaysAgo, *sortByCount)
}
