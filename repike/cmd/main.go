package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jreisinger/gokatas/repike"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: match <regexp> # reads from stdin")
		os.Exit(2)
	}

	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()

	n := 0
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if repike.Match(os.Args[1], line) {
			fmt.Fprintln(output, line)
			n++
		}
	}
	if n == 0 {
		os.Exit(1)
	}
}
