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
	output, err := pmset()
	if err != nil {
		return Status{}, err
	}
	n, err := parse(output)
	if err != nil {
		return Status{}, err
	}
	return Status{ChargedPercent: n}, nil
}

// Run pmset command and return its output.
func pmset() (string, error) {
	output, err := exec.Command("pmset", "-g", "ps").Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

var percent = regexp.MustCompile(`(\d+)%`)

// Parse ChargedPercent from the pmset command output.
func parse(output string) (int, error) {
	matches := percent.FindStringSubmatch(output)
	if len(matches) < 2 {
		return 0, fmt.Errorf("can't parse percent from %q", output)
	}
	n, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, fmt.Errorf("can't convert %v to int", matches[1])
	}
	return n, nil
}
