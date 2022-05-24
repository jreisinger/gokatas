// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
// Adapted from github.com/adonovan/gopl.io/blob/master/ch7/bytecounter.
package main

import (
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	c = 0 // reset the counter
	fmt.Fprintf(&c, "world")
	fmt.Println(c)
}
