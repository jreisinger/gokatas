package main

import (
	"log"
	"net"

	"github.com/jreisinger/gokatas/clock2"
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
		// clock2.HandleConn(conn) // handle one connection at a time
		go clock2.HandleConn(conn) // handle connections concurrently
	}
}
