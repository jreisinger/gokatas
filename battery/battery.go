/*
Package battery gets the MacBook battery status. It shows how to run external
commands and how to run tests conditionally.

What can we test? We don't want to test the external command we call but that:

 1. We execute the `pmset` command with correct arguments.
 2. We correctly parse output to get battery status.

We skip 1. since it's trivial.

Level: intermediate
Topics: integration tests, exec, regexp, tpg-tools
*/
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
	output, err := getPmsetOutput()
	if err != nil {
		return Status{}, err
	}
	status, err := parsePmsetOutput(output)
	if err != nil {
		return Status{}, err
	}
	return status, nil
}

var percentage = regexp.MustCompile(`(\d+)%`)

func parsePmsetOutput(output string) (Status, error) {
	matches := percentage.FindStringSubmatch(output)
	if len(matches) < 2 {
		return Status{}, fmt.Errorf(
			"failed to parse pmset output: %q", output)
	}
	i, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, fmt.Errorf(
			"failed to parse pmset ouput: %v", err)
	}
	return Status{ChargePercent: i}, nil
}

func getPmsetOutput() (string, error) {
	data, err := exec.Command("pmset", "-g", "ps").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}
