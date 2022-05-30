// Areader reads and prints three bytes from Areader.
package main

import (
	"fmt"

	"github.com/jreisinger/gokatas/areader"
)

func main() {
	var r areader.Areader
	p := make([]byte, 3)
	r.Read(p) // NOTE: ignoring error
	fmt.Println(string(p))
}
