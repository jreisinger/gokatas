// Package count3 uses "functional options" pattern to set zero or more options.
// Adapted from https://github.com/bitfield/tpg-tools2/tree/main/count/3
//
// Level: intermediate
// Topics: functional options, tpg-tools
package count3

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct {
	input io.Reader
}

type option func(c *counter) error

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil { // validate option
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func NewCounter(opts ...option) (*counter, error) {
	c := &counter{
		input: os.Stdin,
	}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *counter) Lines() int {
	var lines int
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		lines++
	}
	return lines
}

func Main() {
	c, err := NewCounter()
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Lines())
}
