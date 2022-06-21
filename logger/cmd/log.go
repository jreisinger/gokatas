// Log handles logging correctly. When it's not possible to write logs for some
// reason the whole program will not stop.
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
	for d.problem { // simulate a problem, e.g. disk or network issue
		time.Sleep(time.Second)
	}

	fmt.Print(string(p))
	return len(p), nil
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
				time.Sleep(100 * time.Millisecond)
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
