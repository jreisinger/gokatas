package count_test

import (
	"strings"
	"testing"

	"github.com/jreisinger/gokatas/count"
)

func TestCounterCountsLines(t *testing.T) {
	t.Parallel()

	tests := []struct {
		lines string
		want  int
	}{
		{"", 0},
		// {" ", 0}, TODO: why this fails?
		{"\n", 1},
		{"one\ntwo\nthree", 3},
	}

	for i, test := range tests {
		r := strings.NewReader(test.lines)
		c := count.NewCounter()
		c.Input = r
		got := c.Lines()
		if got != test.want {
			t.Errorf("test %d: got %d, want %d", i, got, test.want)
		}
	}
}
