// Select statement is another way to handle multiple channels. It's like switch
// but each case is a communication.
//
//   - All channels are evaluated.
//   - Blocks until one communication can proceed.
//   - If multiple can proceed, chooses (pseudo-)randomly.
//   - A default case, if present, executes immediately if no channel is ready.
//
// Timeout the conversation when no one speaks for 600 ms.
//
// Level: intermediate
// Topics: select, timeout, fan-in
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(say("Ann"), say("Joe"))
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(time.Millisecond * 600):
			fmt.Println("bye bye")
			return
		}
	}
}

func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { // only one goroutine is needed now
		for {
			select {
			case s := <-c1:
				c <- s
			case s := <-c2:
				c <- s
			}
		}
	}()
	return c
}

func say(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s, %d", msg, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()
	return c
}
