// Log handles logging gracefully; the goroutines that do some important task
// (like sleeping) will not block just because it's not possible to write logs.
//
// Start 10 goroutines each of which will be writing logs to a device. Simulate
// a device problem by pressing Ctrl-C. Press Ctrl-C again to "fix" the problem.
// Ctrl-\ will terminate the program (with a core dump).
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
	// var l log.Logger // the standard logger is a blocking logger
	// l.SetOutput(&d)
	l := logger.New(&d, grs)

	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				l.Write(fmt.Sprintf("log from gr #%d", id))
				doSomething()
			}
		}(i)
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	for {
		<-sigint
		d.problem = !d.problem
	}
}

func doSomething() {
	time.Sleep(time.Second)
}
