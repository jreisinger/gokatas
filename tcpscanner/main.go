// Tcpscanner reports open TCP ports on a host. First create a pool of n workers
// that will do the scanning by connecting to ports. Then send them port numbers
// to try to connect to. Collect the results, 0 means couldn't connect, and
// print them. Adapted from the "Black Hat Go" book.
//
// Topics: concurrency, security
// Level: intermediate
package main

import (
	"fmt"
	"net"
	"sort"
)

const (
	host     = "scanme.nmap.org"
	minPort  = 1
	maxPort  = 1024
	nWorkers = 100
)

func worker(ports, results chan int) {
	for port := range ports {
		addr := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			// Port closed (syn->, <-rst) or
			// filtered (syn->, timeout).
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

	for i := 0; i < nWorkers; i++ {
		go worker(ports, results)
	}

	go func() {
		for i := minPort; i <= maxPort; i++ {
			ports <- i
		}
	}()

	var openports []int

	for i := minPort; i <= maxPort; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	sort.Ints(openports)
	fmt.Println(host, openports)
}
