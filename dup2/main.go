// Dup2 prints the count and text of lines that appear more than once in the
// input. It reads from stdin or from a list of named files. Adapted from
// github.com/adonovan/gopl.io/tree/master/ch1/dup2.
//
// Level: beginner
// Topics: io, maps, bufio.Scanner
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	if len(os.Args[1:]) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range os.Args[1:] {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
