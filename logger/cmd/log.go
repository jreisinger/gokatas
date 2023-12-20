// Log handles logging gracefully. When it's not possible to write logs the
// goroutines will not block.
//
// Start 10 goroutines each of which will be writing logs to a device. Simulate
// a device problem by pressing Ctrl-C. Press Ctrl-C again to "fix" the problem.
// Ctrl-\ will terminate the program (with a core dump).
package main

import (
	"fmt"
	"math/rand"
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

	// Standard logger is a blocking logger.
	// var l log.Logger
	// l.SetOutput(&d)

	l := logger.New(&d, grs)

	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				l.Write(fmt.Sprintf("log from gr %d", id))
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			}
		}(i)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	for {
		<-sigs
		d.problem = !d.problem
	}
}
