// Count ".go" files in fs.FS.
package main

import (
	"fmt"
	"os"

	"github.com/jreisinger/gokatas/countgo"
)

func main() {
	fsys := os.DirFS(os.Args[1])
	fmt.Println(countgo.Files(fsys))
}
