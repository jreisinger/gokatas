// Pingpong shows how (un-buffered) channels are used for communication and
// synchronization between goroutines. When pinger or ponger attempts to send a
// message on the channel, it will wait (block) until printer is ready to
// receive the message. Note that "ping" and "pong" are alternating. This is
// because channels act as first-in-first-out queues (https://go.dev/ref/spec).
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go pinger(ch)
	go ponger(ch)
	go printer(ch)

	var enter string
	fmt.Scanln(&enter)
}

func pinger(ch chan<- string) {
	for {
		ch <- "ping"
	}
}
func ponger(ch chan<- string) {
	for {
		ch <- "pong"
	}
}
func printer(ch <-chan string) {
	for {
		fmt.Printf("%s ", <-ch)
		time.Sleep(time.Second)
	}
}
