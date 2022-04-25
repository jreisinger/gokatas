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
			log.Print(err) // e.g., connection aborted
			continue
		}
		// don't put any code here
		go proxy(conn)
	}
}

func proxy(conn net.Conn) {
	defer conn.Close()

	upstream, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		log.Print(err)
		return
	}
	defer upstream.Close() // to release precious file descriptor

	go io.Copy(upstream, conn) // in this case it's ok not to track goroutine
	io.Copy(conn, upstream)
}
