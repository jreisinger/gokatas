package count2_test

import (
	"strings"
	"testing"

	"github.com/jreisinger/gokatas/count2"
)

func TestLinesCountsLinesInStdin(t *testing.T) {
	t.Parallel()
	r := strings.NewReader("one\ntwo\nthree")
	c, err := count2.NewCounter(
		count2.WithInput(r),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestLinesCountsLinesInCLIargs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := count2.NewCounter(
		count2.WithInputFromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
