// Enum shows how to use enumerated type.
package main

import (
	"fmt"

	"github.com/jreisinger/gokatas/enum"
)

func main() {
	d := enum.North
	fmt.Print(d)
	switch d {
	case enum.North:
		fmt.Println(" goes up.")
	case enum.South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}
}
