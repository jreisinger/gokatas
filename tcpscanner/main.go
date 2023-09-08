// Tcpscanner reports open TCP ports on a host. First create a pool of workers
// that will do the scanning by connecting to ports. Then send them port numbers
// to try to connect to. Collect the results, 0 means couldn't connect, and
// print them. Adapted from the "Black Hat Go" [book].
//
// Topics: concurrency, security
// Level: intermediate
//
// [book]: https://github.com/blackhat-go/bhg/blob/master/ch-2/tcp-scanner-final
package main

import (
	"fmt"
	"net"
	"sort"
)

const host = "scanme.nmap.org"

func worker(ports, results chan int) {
	for port := range ports {
		addr := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			// closed port		syn->, <-rst
			// filtered port	syn->, timeout
			results <- 0
			continue
		}
		conn.Close()
		results <- port
	}
}

func main() {
	in := make(chan int, 100) // can hold 100 items before sender blocks
	out := make(chan int)

	for i := 0; i < cap(in); i++ {
		go worker(in, out)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			in <- i
		}
	}()

	var openports []int

	for i := 1; i <= 1024; i++ {
		port := <-out
		if port != 0 {
			openports = append(openports, port)
		}
	}

	sort.Ints(openports)
	fmt.Println(host, openports)
}
