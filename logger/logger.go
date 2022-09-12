// Package logger uses channels to implement non-blocking logging. Adapted from
// https://youtu.be/zDCKZn4-dck.
//
// Level: advanced
// Topics: concurrency, design
package logger

import (
	"fmt"
	"io"
	"sync"
)

type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

func New(w io.Writer, cap int) *Logger {
	l := Logger{
		ch: make(chan string, cap),
	}

	l.wg.Add(1)
	go func() {
		for v := range l.ch {
			fmt.Fprint(w, v)
		}
		l.wg.Done()
	}()

	return &l
}

func (l *Logger) Stop() {
	close(l.ch)
	l.wg.Wait()
}

func (l *Logger) Println(s string) {
	select {
	case l.ch <- s + "\n":
	default:
		fmt.Println("not logging")
	}
}
