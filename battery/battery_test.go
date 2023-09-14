package battery_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jreisinger/gokatas/battery"
)

func TestParsePmsetOutput_GetsChargePercent(t *testing.T) {
	t.Parallel()
	input, err := os.ReadFile("testdata/pmset.txt")
	if err != nil {
		t.Error(err)
	}
	want := battery.Status{
		ChargePercent: 94,
	}
	got, err := battery.ParsePmsetOutput(string(input))
	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
