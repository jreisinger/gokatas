// You can use channels as a handle on a service. You communicate with the
// service via the channel.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
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
