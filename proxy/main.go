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

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
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
	defer upstream.Close() // to release file descriptor

	go io.Copy(upstream, conn) // here it's ok not to track goroutine
	io.Copy(conn, upstream)
}
