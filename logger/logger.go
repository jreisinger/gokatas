// Package logger uses channels to implement non-blocking
// logging. Adapted from https://youtu.be/zDCKZn4-dck.
//
// Level: advanced
// Topics: design, buffered channels, os/signal
package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type Logger struct {
	logs chan string
	wg   sync.WaitGroup
}

// New creates a logger that will write logs to w. Buf is the size of logs buffer.
func New(w io.Writer, buf int) *Logger {
	// New is sometimes called a factory function. It's useful
	// when you need to initialize one or more fields of a type.
	l := Logger{
		logs: make(chan string, buf),
	}

	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		for log := range l.logs {
			fmt.Fprintln(w, log)
		}
	}()

	return &l
}

// Stop stops accepting logs and waits for logs buffer to be written.
func (l *Logger) Stop() {
	close(l.logs)
	l.wg.Wait()
}

// Write writes the log. If the log buffer is full it prints a warning and exits.
func (l *Logger) Write(log string) {
	select {
	case l.logs <- log:
	default:
		fmt.Fprintln(os.Stderr, "WARN: dropping logs")
	}
}
