// Log handles logging correctly. When it's not possible to write logs for some
// reason the whole program will not block.
//
// Start 10 goroutines each of which will be writing logs to a device. Simulate
// a device problem (e.g. disk or network issue) by pressing Ctrl-C. Press
// Ctrl-C again to "fix" the problem. Ctrl-\ will terminate the program (with a
// core dump).
package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jreisinger/gokatas/logger"
)

type device struct {
	problem bool
}

func (d *device) Write(p []byte) (int, error) {
	for d.problem {
		time.Sleep(time.Second)
	}

	return fmt.Print(string(p))
}

func main() {
	const grs = 10
	var d device

	// Blocking logger.
	// var l log.Logger
	// l.SetOutput(&d)

	// Non-blocking logger.
	l := logger.New(&d, grs)

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
		d.problem = !d.problem
	}
}
