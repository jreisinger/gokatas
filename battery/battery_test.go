package battery

import (
	"os"
	"testing"
)

func TestWeParseChargedPercentFromPmsetOutput(t *testing.T) {
	t.Parallel()
	output, err := os.ReadFile("./testdata/pmset.txt")
	if err != nil {
		t.Fatal(err)
	}
	n, err := parse(string(output))
	if err != nil {
		t.Fatal(err)
	}
	if n != 94 {
		t.Errorf("got %d, want 94", n)
	}
}

func TestPmsetReturnsSomething(t *testing.T) {
	t.Parallel()
	output, err := pmset()
	if err != nil {
		t.Skipf("can't run pmset: %v", err)
	}
	if output == "" {
		t.Errorf("pmset returned nothing")
	}
}
