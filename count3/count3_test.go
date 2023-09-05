package count3_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/jreisinger/gokatas/count3"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"count": count3.Main,
	}))
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestLinesCountsLinesInStdin(t *testing.T) {
	t.Parallel()
	input := strings.NewReader("one\ntwo\nthree")
	c, err := count3.NewCounter(
		count3.WithInput(input),
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
	c, err := count3.NewCounter(
		count3.WithInputFromArgs(args),
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

func TestLinesCountsLinesInStdin_IgnoresEmptyArgs(t *testing.T) {
	t.Parallel()
	args := []string{}
	input := bytes.NewBufferString("1\n2\n3")
	c, err := count3.NewCounter(
		count3.WithInput(input),
		count3.WithInputFromArgs(args),
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
