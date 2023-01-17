// Tcpscanner reports open TCP ports on a host. First create a pool of workers
// that will do the scanning by connecting to ports. Then send them port numbers
// to try to connect to. Collect the results, 0 means couldn't connect, and
// print them. Adapted from the "Black Hat Go" book.
//
// Topics: concurrency, security, scripting
// Level: intermediate
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
	ports := make(chan int)
	results := make(chan int)

	for i := 0; i < 100; i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	var openports []int

	for i := 1; i <= 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	sort.Ints(openports)
	fmt.Println(host, openports)
}
