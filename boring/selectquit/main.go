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
	quit := make(chan string)
	c := boring("blahblah", quit)
	n := rand.Intn(200)
	fmt.Println(n)
	for i := n; i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "bye"
	fmt.Printf("They said: %s. \n", <-quit)
}

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			n := rand.Intn(1e3)
			time.Sleep(time.Duration(n) * time.Millisecond)
			select {
			case c <- fmt.Sprintf("%s,%d", msg, i):
				//do nothing - why?
			case <-quit:
				cleanup() //what does it do?
				quit <- "see you"
				return
			}
		}
	}()
	return c
}
func cleanup() {}
