package count2_test

import (
	"strings"
	"testing"

	"github.com/jreisinger/gokatas/count2"
)

func TestLines_CountsLinesInInput(t *testing.T) {
	t.Parallel()
	c := count2.NewCounter()
	c.Input = strings.NewReader("1\n2\n3\n")
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
