// A channel allows for communication and synchronization between goroutines.
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
	c := make(chan string)
	go say("blah", c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}

func say(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s, %d", msg, i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	}
}
