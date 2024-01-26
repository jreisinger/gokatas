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
	"time"
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
			fmt.Fprintf(w, "%s: %s\n", time.Now().Format(time.TimeOnly), log)
		}
	}()

	return &l
}

func (l *Logger) Write(log string) {
	select {
	case l.logs <- log:
	default:
		fmt.Fprintf(os.Stderr, "%s: dropping logs\n", time.Now().Format(time.TimeOnly))
	}
}
