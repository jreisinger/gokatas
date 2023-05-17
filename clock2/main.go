// Clock2 is a TCP server that periodically writes the time. Adapted from
// github.com/adonovan/gopl.io/tree/master/ch8/clock2.
//
// Level: intermediate
// Topics: networking, TCP server
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil { // e.g., connection aborted
			log.Print(err)
			continue
		}
		go handle(conn) // handle connections concurrently
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil { // e.g., client disconnected
			log.Print(err)
			return
		}
		time.Sleep(time.Second)
	}
}
