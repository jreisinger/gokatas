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
	// var l log.Logger
	// l.SetOutput(&d)
	l := logger.New(&d, grs)

	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				l.Println(fmt.Sprintf("%d: log data", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		<-sigChan
		d.problem = !d.problem
	}
}
