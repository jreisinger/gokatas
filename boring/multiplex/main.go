// Multiplex lets whosoever is ready talk using fanIn function.
//
// Level: intermediate
// Topics: goroutines, channels
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Ann"), boring("Joe"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func fanIn(input1, input2 <-chan string) <-chan string {

	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input1
		}
	}()
	return c
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s, %d", msg, i)
			r := rand.Intn(2e3)
			time.Sleep(time.Duration(time.Duration(r)) * time.Millisecond)
		}
	}()
	return c
}
