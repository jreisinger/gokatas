package battery

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWeGetNonEmptyOutputWhenWeRunPmset(t *testing.T) {
	t.Parallel()
	output, err := runPmset()
	if err != nil {
		t.Skipf("unable to run 'pmset' command: %v", err)
	}
	if output == "" {
		t.Errorf("we got no output from pmset")
	}
}

func TestWeParseOutChargedPercentFromPmsetOutput(t *testing.T) {
	t.Parallel()
	output, err := os.ReadFile("./testdata/pmset.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := Status{
		ChargedPercent: 94,
	}
	got, err := parsePmset(string(output))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
