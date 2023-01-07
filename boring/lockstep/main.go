// Lockstep makes Ann and Joe talk in lockstep (one after another). You can use
// channels as a handle on a service.
//
// Level: beginner
// Topics: goroutines, channels
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ann := say("Ann")
	joe := say("Joe")
	for i := 0; i < 5; i++ {
		fmt.Println(<-ann)
		fmt.Println(<-joe)
	}
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
