// Counter creates an integer counter, increments it, prints it and resets it.
package main

import (
	"fmt"

	"github.com/jreisinger/gokatas/counter"
)

func main() {
	c := new(counter.Counter)
	c.Increment()
	fmt.Println(c.N())
	c.Reset()
}
