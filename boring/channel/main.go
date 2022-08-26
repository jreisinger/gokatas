// A channel allows for communication and synchronization between goroutines.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan string)
	go boring("blah", c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s, %d", msg, i)
		n := rand.Intn(1e3)
		time.Sleep(time.Millisecond * time.Duration(n))
	}
}
