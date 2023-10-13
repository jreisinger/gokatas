// Package battery gets MacBook battery status. It shows how to run external
// commands and how to parse their output.
//
// Level: intermediate
// Topics: exec, regexp, tpg-tools
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
	output, err := runPmset()
	if err != nil {
		return Status{}, err
	}
	return parsePmsetOutput(output)
}

func runPmset() (string, error) {
	data, err := exec.Command("pmset", "-g", "ps").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

var percentage = regexp.MustCompile(`(\d+)%`)

func parsePmsetOutput(output string) (Status, error) {
	matches := percentage.FindStringSubmatch(output)
	if len(matches) != 2 {
		return Status{}, fmt.Errorf("can't parse %q", output)
	}
	perc, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, fmt.Errorf("can't parse %q: %v", output, err)
	}
	return Status{ChargePercent: perc}, nil
}
