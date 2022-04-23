package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jreisinger/gokatas"
)

var katasFile = flag.String("f", "katas.md", "file with katas you've done")
var showAll = flag.Bool("a", false, "show all katas (default is those last done within two weeks)")
var countSort = flag.Bool("c", false, "sort katas by count (default is by last done)")

func main() {
	flag.Parse()
	if err := gokatas.PrintStats(*katasFile, *showAll, *countSort); err != nil {
		fmt.Fprintf(os.Stderr, "gokatas: %v\n", err)
		os.Exit(1)
	}

	// var katas []gokatas.Kata
	// katas, err := gokatas.ParseFile(*katasFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// stats := gokatas.GetStats(katas)
	// gokatas.PrintStats(stats, showAll, countSort)
}
