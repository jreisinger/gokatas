// Netcat2 is a read/write TCP client. Adapted from
// github.com/adonovan/gopl.io/tree/master/ch8/netcat2.
//
// Level: intermediate
// Topics: networking, TCP client, read-write
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
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
