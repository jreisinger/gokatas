// Netcat1 is a read-only TCP client. Adapted from
// github.com/adonovan/gopl.io/tree/master/ch8/netcat1.
//
// Level: intermediate
// Topics: networking
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1362")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
