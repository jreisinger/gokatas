package count5_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/jreisinger/gokatas/count5"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"count": count5.Main,
	}))
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestLines_CountsLinesInInput(t *testing.T) {
	t.Parallel()
	input := strings.NewReader("one\ntwo\nthree")
	c, err := count5.NewCounter(
		count5.WithInput(input),
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

func TestLines_CountsLinesInCLIargs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := count5.NewCounter(
		count5.WithInputFromArgs(args),
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

func TestLines_CountsLinesInStdin_IgnoresEmptyArgs(t *testing.T) {
	t.Parallel()
	args := []string{}
	input := bytes.NewBufferString("1\n2\n3")
	c, err := count5.NewCounter(
		count5.WithInput(input),
		count5.WithInputFromArgs(args),
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
