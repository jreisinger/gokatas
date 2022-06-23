// Proxy mediates TCP traffic between client and upstream. Adapted from
// youtu.be/J4J-A9tcjcA.
package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for { // accept loop
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// Don't put any blocking code here! Because it could block your
		// accept loop and thus the whole server :-/. E.g., a client
		// could open a connection and not send anything.
		go proxy(conn)
	}
}

func proxy(conn net.Conn) {
	defer conn.Close() // to release precious file descriptor

	upstream, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		log.Print(err)
		return
	}
	defer upstream.Close()

	go io.Copy(upstream, conn) // in this case it's ok not track goroutine
	io.Copy(conn, upstream)
}
