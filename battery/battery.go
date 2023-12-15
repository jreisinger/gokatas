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
	ChargedPercent int
}

func GetStatus() (Status, error) {
	output, err := runPmset()
	if err != nil {
		return Status{}, err
	}
	status, err := parsePmset(output)
	if err != nil {
		return Status{}, err
	}
	return status, nil
}

func runPmset() (string, error) {
	output, err := exec.Command("/usr/bin/pmset", "-g", "ps").Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

var percent = regexp.MustCompile(`(\d+)%`)

func parsePmset(output string) (Status, error) {
	matches := percent.FindStringSubmatch(output)
	if len(matches) < 2 {
		return Status{}, fmt.Errorf("no percent match in %q", output)
	}
	n, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, fmt.Errorf("can't convert %q to int", matches[1])
	}
	return Status{ChargedPercent: n}, nil
}
