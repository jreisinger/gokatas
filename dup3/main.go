// Dup3 prints the count and text of lines that appear more than once in the
// named input files. It reads in entire file at once ("slurp" mode). Adapted
// from github.com/adonovan/gopl.io/tree/master/ch1/dup3.
//
// os.ReadFile (dup3) and bufio.Scanner (dup2) use the Read and Write methods
// of *os.File under covers.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, arg := range os.Args[1:] {
		b, err := os.ReadFile(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(b), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
