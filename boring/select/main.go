/*
Select statement is another way to handle multiple channels. It's like switch
but each case is a communication:

 - All channels are evaluated.
 - Blocks until one communication can proceed.
 - If multiple can proceed, chooses pseudo-randomly.
 - A default case, if present, executes immediately if no channel is ready.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Ann"), boring("Joe"))
	// you could also timeout the whole conversation see
	// https://go.dev/talks/2012/concurrency.slide#36
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("You guys are too slow.")
			return
		}
	}
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { // only one goroutine is needed now
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)

			n := rand.Intn(1e3)
			time.Sleep(time.Duration(n) * time.Millisecond)
		}
	}()
	return c
}
