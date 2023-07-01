// Areader reads three bytes from A into a slice.
package main

import (
	"fmt"

	"github.com/jreisinger/gokatas/areader"
)

func main() {
	var a areader.A
	p := make([]byte, 3)
	a.Read(p)
	fmt.Println(string(p))
}
