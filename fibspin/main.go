// Fibspin calculates 46th number from the Fibonacci sequence. It uses a slow
// recursive algorithm. While the calculation is being done a spinner is
// displayed. Adapted from github.com/adonovan/gopl.io/tree/master/ch8/spinner.
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(time.Millisecond * 100)
	fmt.Printf("\rfib(45) = %d\n", fib(45))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `\|/-` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}
