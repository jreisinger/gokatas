package main

import (
	"fmt"

	"github.com/jreisinger/gokatas/count"
)

func main() {
	c := count.NewCounter()
	fmt.Println(c.Lines())
}
