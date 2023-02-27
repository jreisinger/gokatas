// Bytecounter shows how to use interfaces. It implements a concrete type named
// ByteCounter whose Write method counts bytes. Since ByteCounter safisfies the
// io.Writer interface (an abstract type), we can pass it to fmt.Fprint.
//
// Adapted from github.com/adonovan/gopl.io/blob/master/ch7/bytecounter.
//
// Level: beginner
// Topics: interfaces, io.Writer
package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	c += 10

	fmt.Fprint(&c, "worlddddd")
	fmt.Println(c)
}
