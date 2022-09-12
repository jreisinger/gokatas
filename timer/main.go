// Timer times out a goroutine that emits Go pointer types over a channel.
// Based on "Go Programming Basics" video from John Graham-Cumming.
//
// Level: intermediate
// Topics: time.Timer, concurrency
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go pointerTypes(ch)
	for pt := range ch {
		fmt.Println(pt)
		time.Sleep(time.Second)
	}
}

func pointerTypes(ch chan string) {
	defer close(ch)
	ptypes := []string{"slices", "maps", "functions", "channels", "interfaces"}
	var i int
	t := time.NewTimer(time.Second * 5)
	for {
		select {
		case ch <- ptypes[i]:
			i++
			if i == len(ptypes) {
				i = 0
			}
		case <-t.C:
			return
		}
	}
}
