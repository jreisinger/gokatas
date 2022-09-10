// Boring listens for a while to a boring person talking in a goroutine.
// Topics: concurrency
package main

import (
	"fmt"
	"time"
)

func main() {
	go boring("blah") // analogous to the & on the end of a shell command
	time.Sleep(time.Second * 5)
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Printf("%s, %d\n", msg, i)
		time.Sleep(time.Second)
	}
}
