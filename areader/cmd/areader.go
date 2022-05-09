// Areader reads and prints three bytes from Areader.
package main

import (
	"fmt"

	"github.com/jreisinger/gokatas/areader"
)

func main() {
	p := make([]byte, 3)
	r := areader.Areader{}
	r.Read(p) // NOTE: ignoring error
	fmt.Println(string(p))
}
