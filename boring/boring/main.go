// Boring listens for a while to a boring person talking in a goroutine.
//
// Based on Go Concurrency Patterns by Rob Pibe (2012):
//
//	Slides	https://talks.golang.org/2012/concurrency.slide
//	Code	https://talks.golang.org/2012/concurrency/support
//	Video	https://www.youtube.com/watch?v=f6kdp27TYZs
//
// Level: beginner
// Topics: goroutines
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
