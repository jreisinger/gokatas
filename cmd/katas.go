package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jreisinger/gokatas"
)

var showAll = flag.Bool("a", false, "show all katas (default is those last done within two weeks)")
var countSort = flag.Bool("c", false, "sort katas by count (default is by last done)")

func main() {
	flag.Parse()
	if err := gokatas.PrintStats(*showAll, *countSort); err != nil {
		fmt.Fprintf(os.Stderr, "gokatas: %v\n", err)
		os.Exit(1)
	}
}
