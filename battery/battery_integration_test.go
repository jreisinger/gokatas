//go:build integration

// go test -tags=integration
package battery_test

import (
	"bytes"
	"os/exec"
	"testing"

	"github.com/jreisinger/gokatas/battery"
)

func TestGetPmsetOutput_CapturesCmdOutput(t *testing.T) {
	t.Parallel()
	data, err := exec.Command("/usr/bin/pmset", "-g", "ps").CombinedOutput()
	if err != nil {
		t.Skipf("unable to run 'pmset' command: %v", err)
	}
	if !bytes.Contains(data, []byte("InternalBattery")) {
		t.Skip("no battery fitted")
	}
	output, err := battery.GetPmsetOutput()
	if err != nil {
		t.Fatal(err)
	}
	status, err := battery.ParsePmsetOutput(output)
	if err != nil {
		t.Fatal(err)
	}
	// Will be printed only when test fails or -v is used.
	t.Logf("Charge: %d%%", status.ChargePercent)
}
