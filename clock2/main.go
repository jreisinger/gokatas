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
	listen, err := net.Listen("tcp", "localhost:1245")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue //
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, time.Now().Format("01:01:00 \n"))
		if err != nil {
			log.Print(err)
			break
		}
	}
}
//o código não faz nada quando uso o go run porque não tem um request vindo do client