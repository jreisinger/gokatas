// Package count2 is like count but it allows to set zero or more options on a
// counter. It uses "functional options" pattern to do so.
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
	ouput io.Writer
}

type option func(*Counter) error

func WithInput(input io.Reader) option {
	return func(c *Counter) error {
		if input == nil { // validate option
			return errors.New("nil input reader")
		}
		c.input = input
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(c *Counter) error {
		if len(args) < 1 {
			return nil
		}
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		c.input = f
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *Counter) error {
		if output == nil { // validate option
			return errors.New("nil output writer")
		}
		c.ouput = output
		return nil
	}
}

func NewCounter(opts ...option) (*Counter, error) {
	c := &Counter{
		input: os.Stdin,
		ouput: os.Stdout,
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
	input := bufio.NewScanner(c.input)
	for input.Scan() {
		lines++
	}
	return lines
}
