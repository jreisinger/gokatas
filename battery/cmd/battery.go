package main

import (
	"fmt"
	"os"

	"github.com/jreisinger/gokatas/battery"
)

func main() {
	status, err := battery.GetStatus()
	if err != nil {
		fmt.Fprintf(os.Stderr, "battery: getting status: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("charged %d%%\n", status.ChargedPercent)
}
