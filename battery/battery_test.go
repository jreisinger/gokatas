package battery

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParsePmsetOutput(t *testing.T) {
	t.Parallel()
	want := Status{
		ChargedPercent: 94,
	}
	data, err := os.ReadFile("testdata/pmset.txt")
	if err != nil {
		t.Fatal(err)
	}
	got, err := parsePmsetOutput(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
