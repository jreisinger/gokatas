//go:build integration

/*
Unit tests
- test functions work as expected in isolation

Integration tests
- test assumptions about how external dependencies work are still true
- can be slower because we run them less often

go test -tags=integration
*/

package battery

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestGetPmsetOutput_CapturesCmdOutput(t *testing.T) {
	t.Parallel()
	data, err := exec.Command("/usr/bin/pmset", "-g", "ps").CombinedOutput()
	if err != nil {
		t.Skipf("failed to execute 'pmset' command: %v", err)
	}
	if !bytes.Contains(data, []byte("InternalBattery")) {
		t.Skip("no battery fitted")
	}
	output, err := getPmsetOutput()
	if err != nil {
		t.Fatal(err)
	}
	status, err := parsePmsetOutput(output)
	if err != nil {
		t.Fatal(err)
	}
	// Will be printed only when test fails or -v is used.
	t.Logf("battery charge: %d%%", status.ChargePercent)
}
