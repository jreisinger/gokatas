package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jreisinger/gokatas/older"
)

var (
	in   = flag.String("in", ".", "find files in `path`")
	than = flag.Duration("than", time.Hour*24*30, "find files older than `duration`")
)

func main() {
	flag.Parse()
	fsys := os.DirFS(*in)
	files := older.Files(fsys, *than)
	for _, f := range files {
		fmt.Println(f)
	}
}
