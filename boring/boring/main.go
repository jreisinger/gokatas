// Package boring and its subpackages hold various Go concurrency patterns in
// the form of boring conversations. This is first of them. It's not an honest
// one because there is no communication.
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
	"time"
)

func main() {
	go boring("blah") // analogous to the & on the end of a shell command
	time.Sleep(time.Second * 10)
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Printf("%s, %d\n", msg, i)
		time.Sleep(time.Second)
	}
}
