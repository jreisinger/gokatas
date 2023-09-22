package count3_test

import (
	"strings"
	"testing"

	"github.com/jreisinger/gokatas/count3"
)

func TestLines_CountsLinesInInput(t *testing.T) {
	t.Parallel()
	input := strings.NewReader("1\n2\n3")
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
