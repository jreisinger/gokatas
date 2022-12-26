// Lockstep makes Ann and Joe talk in lockstep. You can use channels as a handle
// on a service.
//
// Level: beginner
// Topics: goroutines, channels
package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	ann := boring("Ann")
	joe := boring("Joe")

	for i:=0;i<=10;i++{
		fmt.Println(<-ann)
		fmt.Println(<-joe)
	}

}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)

			rnd := rand.Intn(1e3)
			time.Sleep(time.Duration(rnd) * time.Millisecond)

		}

	}()
	return c
}
