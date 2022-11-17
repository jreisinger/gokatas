// Generator is a function that returns a channel. We launch the goroutine from
// inside the generator.
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
	c:= boring("esse é meu texto chato")

	for i:=0;i<8;i++{
		fmt.Println(<- c)
	}
}

func boring(msg string) <-chan string {
	canal := make(chan string)

	go func() {
		for i := 0; ; i++ {
			canal <- fmt.Sprintf("%s, %d", msg, i)
			r := rand.Intn(1e2)
			time.Sleep(time.Millisecond * time.Duration(r))

		}
	}()
	return canal
}
