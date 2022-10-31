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
	go boring("blah", c)
	for i := 0; i <= 5; i++ {
		fmt.Println(<- c)
	}
}

func boring(s string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s , %d \n", s, i)
		n := rand.Intn(1e3)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

//criar uma func boring que recebe uma str e um chan de param
//chamar boring na main
