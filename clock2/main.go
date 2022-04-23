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
	addr := net.JoinHostPort("localhost", "1362")
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening at %s", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		// HandleConn(conn) // handle one connection at a time
		go HandleConn(conn) // handle connections concurrently
	}
}

func HandleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(time.Second)
	}
}
