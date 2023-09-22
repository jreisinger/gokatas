// Package count2 counts lines in input. The input defaults to STDIN. Write this
// package with test-first approach (TDD). All non-trivial code should be inside
// the package count2, not main. Adapted from
// https://github.com/bitfield/tpg-tools2/tree/main/count/2
//
// If you want to climb a mountain, begin at the top. -- Zen saying
//
// Level: intermediate
// Topics: TDD, default options, tpg-tools
package count2

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type counter struct {
	Input io.Reader
}

func NewCounter() *counter {
	return &counter{Input: os.Stdin}
}

func (c *counter) Lines() int {
	var lines int
	s := bufio.NewScanner(c.Input)
	for s.Scan() {
		lines++
	}
	return lines
}

func Main() {
	fmt.Println(NewCounter().Lines())
}
