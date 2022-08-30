// Use quit channel to stop the conversation. Also wait for them to tell us
// they're done talking.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan string)
	c := boring("blah", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "bye"
	fmt.Printf("they say: %q\n", <-quit)
}

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			n := rand.Intn(1e3)
			time.Sleep(time.Millisecond * time.Duration(n))
			select {
			case c <- fmt.Sprintf("%s, %d", msg, i):
				// do nothing
			case <-quit:
				cleanup()
				quit <- "see you"
				return
			}
		}
	}()
	return c
}

func cleanup() {}
