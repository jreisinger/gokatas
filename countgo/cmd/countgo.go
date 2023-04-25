// Countgo counts files with .go extension in supplied directory or in ".".
package main

import (
	"fmt"
	"os"

	"github.com/jreisinger/gokatas/countgo"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	fsys := os.DirFS(dir)
	n := countgo.Files(fsys)
	fmt.Printf("%d Go files in %q\n", n, dir)
}
