// Package z takes lines from stdin, makes them into tasks that get processed
// and printed. To use it implement a Factory and a Task. Then call Run() on
// your factory. Is uses concurrency to run fast and interfaces and composition
// to be generic and simple. See https://youtu.be/woCg2zaIVzQ for more.
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
	in := make(chan Task)
	out := make(chan Task)
	var wg sync.WaitGroup

	// Read lines from stdin and stuff them into in channel.
	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- f.Make(s.Text())
		}
		if err := s.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "z: reading from STDIN: %v", err)
		}
		close(in)
		wg.Done()
	}()

	// Read from in channel, do the work, and write results to out channel.
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.Process()
				out <- t
			}
			wg.Done()
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
