package main

import (
	"fmt"

	"github.com/jreisinger/gokatas/counter"
)

func main() {
	c := new(counter.Counter)
	c.Inc()
	fmt.Println(c.Get())
}
