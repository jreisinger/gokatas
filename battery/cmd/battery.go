package main

import (
	"fmt"
	"os"

	"github.com/jreisinger/gokatas/battery"
)

func main() {
	status, err := battery.GetStatus()
	if err != nil {
		fmt.Fprintf(os.Stderr, "battery: can't get status: %v", err)
		os.Exit(1)
	}
	fmt.Printf("battery %d%% charged\n", status.ChargePercent)
}
