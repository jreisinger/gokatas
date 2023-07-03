// Multiplex lets talk Ann or Joe (whosever is ready) using fanIn function.
//
// Level: intermediate
// Topics: goroutines, channels, fan-in
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(say("Ann"), say("Joe"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-c1
		}
	}()
	go func() {
		for {
			c <- <-c2
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
