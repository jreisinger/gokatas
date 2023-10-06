package battery

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_parsePmsetOutput(t *testing.T) {
	t.Parallel()
	want := Status{
		ChargePercent: 94,
	}
	output, err := os.ReadFile("testdata/pmset.txt")
	if err != nil {
		t.Fatal(err)
	}
	got, err := parsePmsetOutput(string(output))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
