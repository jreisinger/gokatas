// Z takes lines from STDIN, makes them into tasks that get processed and
// printed. To use it implement a factory and a task. Then call run() on your
// factory. Is uses concurrency to run fast and interfaces and composition to
// be generic and simple. See https://youtu.be/woCg2zaIVzQ for more.
package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type factory interface {
	make(line string) task
}

type task interface {
	process()
	print()
}

func run(f factory) {
	var wg sync.WaitGroup

	in := make(chan task)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- f.make(s.Text())
		}
		if s.Err() != nil {
			fmt.Fprintf(os.Stderr, "z: reading from STDIN: %v", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan task)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.process()
				out <- t
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.print()
	}
}

func main() {
	// run(&myFactory{})
}
