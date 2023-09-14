// Package battery gets the MacBook battery status. I shows how to run external
// commands and how to run tests conditionally.
//
// Level: intermediate
// Topics: integration tests, exec, regexp, tpg-tools
package battery

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Status struct {
	ChargePercent int
}

func GetStatus() (Status, error) {
	output, err := GetPmsetOutput()
	if err != nil {
		return Status{}, err
	}
	return ParsePmsetOutput(output)
}

var pmsetOutput = regexp.MustCompile(`(\d+)%`)

func ParsePmsetOutput(output string) (Status, error) {
	matches := pmsetOutput.FindStringSubmatch(output)
	if len(matches) != 2 {
		return Status{}, fmt.Errorf("failed to parse pmset output: %q", output)
	}
	charge, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, fmt.Errorf("failed to parse charge percentage: %q", matches[1])
	}
	return Status{ChargePercent: charge}, nil
}

func GetPmsetOutput() (string, error) {
	data, err := exec.Command("/usr/bin/pmset", "-g", "ps").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
