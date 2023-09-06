// Package count3 accepts input also from one or more files supplied as command
// line arguments. It also tests the script using testscript.
//
// Level: advanced
// Topics: testscript, scripting
package count3

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Counter struct {
	files []io.Reader
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
		c.files = make([]io.Reader, len(args))
		for i, path := range args {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			c.files[i] = f
		}
		c.input = io.MultiReader(c.files...)
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
	for _, f := range c.files {
		f.(io.Closer).Close()
	}
	return lines
}

func Main() int {
	c, err := NewCounter(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Println(c.Lines())
	return 0
}
