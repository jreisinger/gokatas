package main

import (
	"fmt"

	"github.com/jreisinger/gokatas/bytecounter"
)

func main() {
	var c bytecounter.Bytecounter
	c.Write([]byte("hello"))
	c = 0 // reset the counter
	fmt.Fprint(&c, "world")
	fmt.Println(c)
}
