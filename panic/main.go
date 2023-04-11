/*
Panic describes and shows how panic and defer work.

Panics come from:
 1. runtime (e.g. out-of-bounds array access or nil pointer dereference)
 2. built-in panic function (for grave errors or impossible situations)

"Expected" errors (from incorrect input, misconfiguration, or failing I/O)
should be handled using error values.

During panic
 1. normal execution stops
 2. all deferred function calls in that goroutine are executed in LIFO order
 3. the program crashes with panic value (usually an error message) + stack trace for each goroutine

Taken from: https://github.com/adonovan/gopl.io/blob/master/ch5/defer1/defer.go

Level: beginner
Topics: panic, defer
*/
package main

import "fmt"

func main() {
	f(3)
}

func f(n int) {
	fmt.Printf("f(%d)\n", n-0/n) // panics when n == 0
	defer fmt.Printf("defer %d\n", n)
	f(n - 1)
}
