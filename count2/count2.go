// Package count2 uses "functional options" pattern to set zero or more options.
//
// Level: intermediate
// Topics: functional options, scripting
package count2

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type Counter struct {
	input io.Reader
}

type option func(c *Counter) error

func WithInput(input io.Reader) option {
	return func(c *Counter) error {
		if input == nil { // validate option
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func NewCounter(opts ...option) (*Counter, error) {
	c := &Counter{
		input: os.Stdin,
	}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Counter) Lines() int {
	var lines int
	scanner := bufio.NewScanner(c.input)
	for scanner.Scan() {
		lines++
	}
	return lines
}
