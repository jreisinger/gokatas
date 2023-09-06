package count_test

import (
	"strings"
	"testing"

	"github.com/jreisinger/gokatas/count"
)

func TestLinesCounting(t *testing.T) {
	r := strings.NewReader("1\n2\n3")
	c := count.NewCounter()
	c.Input = r
	got := c.Lines()
	want := 3
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
