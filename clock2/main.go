// Clock2 is a TCP server that periodically writes the time. Adapted from
// github.com/adonovan/gopl.io/tree/master/ch8/clock2.
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:1362")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		// handle(conn) // handle only one connection at a time
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(time.Second)
	}
}
