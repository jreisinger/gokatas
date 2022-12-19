// Log handles logging correctly. When it's not possible to write logs for some
// reason the whole program will not block.
//
// Start 10 goroutines each of which will be writing logs to a log collector.
// Simulate a log collector problem by pressing Ctrl-C. Press Ctrl-C again to
// "fix" the problem. Ctrl-\ will terminate the program (with a core dump).
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jreisinger/gokatas/logger"
)

type logCollector struct {
	problem bool
}

func (c *logCollector) Write(p []byte) (int, error) {
	for c.problem {
		time.Sleep(time.Second)
	}

	return fmt.Print(string(p))
}

func main() {
	const grs = 10
	var c logCollector

	// Blocking logger.
	// var l log.Logger
	// l.SetOutput(&d)

	// Non-blocking logger.
	l := logger.New(&c, grs)

	for i := 0; i < grs; i++ {
		go func(i int) {
			for {
				l.Println(fmt.Sprintf("gr %d pretending work", i))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	for {
		<-sig
		c.problem = !c.problem
	}
}
