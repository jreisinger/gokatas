// Netcat2 is a read/write TCP client. Adapted from
// github.com/adonovan/gopl.io/tree/master/ch8/netcat2.
//
// Level: intermediate
// Topics: networking, TCP client, read-write
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Fprintf(os.Stderr, "netcat2: supply host and port\n")
		os.Exit(1)
	}
	host, port := os.Args[1], os.Args[2]
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdout)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
