// Areader reads three bytes from Areader into a slice.
package main

import (
	"fmt"

	"github.com/jreisinger/gokatas/areader"
)

func main() {
	var a areader.Areader
	p := make([]byte, 3)
	a.Read(p) // NOTE: ignoring potential error
	fmt.Println(string(p))
}
