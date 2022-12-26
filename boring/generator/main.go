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
c := boring("valeu, natalina")

for i:=0;i<10;i++{
	fmt.Println(<-c)
}
}

func boring(msg string) <-chan string{
	c := make(chan string)

	go func(){
		for i:=0;;i++{
			c <- fmt.Sprintf("%s %d", msg, i)
			r:= rand.Intn(1e3)
			time.Sleep(time.Duration(r) * time.Millisecond)
		}
	}()
	return c
}