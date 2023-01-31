// Bytecounter shows how to use interfaces. It implements a concrete type named
// ByteCounter whose Write method counts bytes before discarding them. Since
// ByteCounter satisfies the io.Writer interface (an abstract type), we can pass
// it to fmt.Fprint.
//
// Adapted from github.com/adonovan/gopl.io/blob/master/ch7/bytecounter.
//
// Level: beginner
// Topics: interfaces, io.Writer
package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // must explicitly convert int to ByteCounter
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello")) // c == 5
	c = 0                    // reset the counter
	fmt.Fprint(&c, "world")  // c == 5
}
