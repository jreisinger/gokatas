package main

import (
	"fmt"
	"os"

	"github.com/jreisinger/gokatas/battery"
)

func main() {
	status, err := battery.GetStatus()
	if err != nil {
		fmt.Fprintf(os.Stderr, "battery: couldn't get status: %v", err)
	}
	fmt.Printf("battery %d%% charged\n", status.ChargePercent)
}
