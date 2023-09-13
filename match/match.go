/*
Package match searches its input for lines containing a given string, and print
them to output. It shows how allow for zero or more options that have default
values.

Level: intermediate
Topics: options, defaults
*/
package match

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
)

type Matcher struct {
	Input  io.Reader
	Output io.Writer
}

type option func(*Matcher) error

func NewMatcher(opts ...option) (*Matcher, error) {
	m := &Matcher{Input: os.Stdin, Output: os.Stdout}
	for _, opt := range opts {
		err := opt(m)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

func WithInput(input io.Reader) option {
	return func(m *Matcher) error {
		if input == nil {
			return errors.New("nil input reader")
		}
		m.Input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(m *Matcher) error {
		if output == nil {
			return errors.New("nil output reader")
		}
		m.Output = output
		return nil
	}
}

func (m *Matcher) PrintLines(pattern *regexp.Regexp) {
	var matched []string
	s := bufio.NewScanner(m.Input)
	for s.Scan() {
		if pattern.MatchString(s.Text()) {
			matched = append(matched, s.Text())
		}
	}
	for _, match := range matched {
		fmt.Fprintln(m.Output, match)
	}
}
