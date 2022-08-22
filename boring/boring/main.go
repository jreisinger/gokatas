// We listen for a while to a boring person talking "in the background".
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go boring("blah") // analogous to the & on the end of a shell command
	time.Sleep(2 * time.Second)
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Printf("%s, %d\n", msg, i)

		n := rand.Intn(1e3)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}
