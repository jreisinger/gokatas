// Lockstep makes Ann and Joe talk in lockstep. You can use channels as a handle
// on a service.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ann := boring("Ann")
	joe := boring("Joe")
	for i := 0; i < 5; i++ {
		fmt.Println(<-ann)
		fmt.Println(<-joe)
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s, %d", msg, i)

			n := rand.Intn(1e3)
			time.Sleep(time.Duration(n) * time.Millisecond)
		}
	}()
	return c
}
