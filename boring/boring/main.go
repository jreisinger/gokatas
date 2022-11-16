// Boring contains various Go concurrency patterns in the form of boring
// conversations. This is first of them. It's not an honest one because there is
// no communication between the main and the boring goroutine.
//
// Based on Go Concurrency Patterns by Rob Pike (2012):
//
//	Slides	https://talks.golang.org/2012/concurrency.slide
//	Code	https://talks.golang.org/2012/concurrency/support
//	Video	https://www.youtube.com/watch?v=f6kdp27TYZs
//
// Level: beginner
// Topics: concurrency, design
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go boring("blah") // independently executing function
	time.Sleep(time.Second * 5)
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Printf("%s, %d\n", msg, i)
		n := rand.Intn(1e3)
		time.Sleep(time.Millisecond * time.Duration(n))
	}
}
