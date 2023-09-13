package match_test

import (
	"bytes"
	"regexp"
	"strings"
	"testing"

	"github.com/jreisinger/gokatas/match"
)

func TestMatchesLines(t *testing.T) {
	input := strings.NewReader("hello world\nwhat's going on here")
	output := new(bytes.Buffer)
	m, err := match.NewMatcher(
		match.WithInput(input),
		match.WithOutput(output),
	)
	if err != nil {
		t.Fatal(err)
	}

	rx := regexp.MustCompile("hello")
	m.PrintLines(rx)
	if output.String() != "hello world\n" {
		t.Errorf("got %q, wanted %q", output.String(), "hello world\n")
	}
}
