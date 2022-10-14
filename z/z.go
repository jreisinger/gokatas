// Package z takes lines from stdin, makes them into tasks that get processed
// and printed. Is uses concurrency to run fast and interfaces and composition
// to be generic and simple. It is a generalized version of lookup-v1.
//
// To use it implement a Factory and a Task. Then call Run() on your factory.
// See lookup-v2/lookup.go for an example.
//
// See https://youtu.be/woCg2zaIVzQ for more.
package z

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type Factory interface {
	Make(line string) Task
}

type Task interface {
	Process()
	Print()
}

func Run(f Factory) {
	var wg sync.WaitGroup
	in := make(chan Task)

	// Read lines from stdin and stuff them into in channel.
	wg.Add(1)
	go func() {
		defer wg.Done()
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- f.Make(s.Text())
		}
		if s.Err() != nil {
			fmt.Fprintf(os.Stderr, "reading STDIN: %v", s.Err())
		}
		close(in)
	}()

	out := make(chan Task)

	// Read from in channel, do the work, and write results to out channel.
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range in {
				t.Process()
				out <- t
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	// Write results from out channel to stdout.
	for t := range out {
		t.Print()
	}
}
