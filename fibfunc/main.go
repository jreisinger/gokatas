// Fibfunc implements a function that returns a function (a closure) generating
// Fibonacci sequence. Then it prints the first 46 numbers from the sequence.
// Adapted from tour.golang.org/moretypes/26.
package main

import "fmt"

func main() {
	f := fib()
	for i := 0; i <= 45; i++ {
		fmt.Printf("fib(%d) = %d\n", i, f())
	}
}

func fib() func() int {
	a, b := 1, 0
	return func() int {
		a, b = b, a+b
		return a
	}
}
