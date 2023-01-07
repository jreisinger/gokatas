// Use quit channel to stop the conversation. Also wait for them to tell us
// they're done talking.
//
// Level: intermediate
// Topics: goroutines, channels, select
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	quit := make(chan string)
	c := say("blah", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "bye"
	fmt.Printf("They said: %q.\n", <-quit)
}

func say(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s, %d", msg, i):
			case <-quit:
				cleanup()
				quit <- "see you"
				return
			}
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()
	return c
}

func cleanup() {}
