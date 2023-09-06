// Package count counts lines in input. The input defaults to STDIN. Write this
// package with test-first approach (TDD). All non-trivial code should be inside
// the package count, not main. Adapted from
// https://bitfieldconsulting.com/books/tools
//
// If you want to climb a mountain, begin at the top. -- Zen saying
//
// Level: intermediate
// Topics: TDD, default options, scripting
package count

import (
	"bufio"
	"io"
	"os"
)

type Counter struct {
	Input io.Reader
}

func NewCounter() Counter {
	return Counter{Input: os.Stdin}
}

func (c Counter) Lines() int {
	var lines int
	scanner := bufio.NewScanner(c.Input)
	for scanner.Scan() {
		lines++
	}
	return lines
}
