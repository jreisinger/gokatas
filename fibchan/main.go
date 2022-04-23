// Fibchan implements a function that returns a channel generating Fibonacci
// sequence (rosettacode.org/wiki/Fibonacci_sequence). Then it prints first 46
// numbers from the sequence.
package main

import "fmt"

func main() {
	c := fib()
	for i := 0; i <= 45; i++ {
		fmt.Printf("fib(%d) = %d\n", i, <-c)
	}
}

func fib() <-chan int {
	c := make(chan int)
	go func() {
		a, b := 0, 1
		for {
			c <- a
			a, b = b, a+b
		}
	}()
	return c
}
