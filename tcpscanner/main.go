// Tcpscanner reports open TCP ports on a host. First create a pool of 100
// workers that will do the scanning by connecting to ports. Then send them
// port numbers to try to connect to. Collect the results, 0 means couldn't
// connect, and print them. Adapted from the "Black Hat Go" book.
//	go run tcpscanner/main.go scanme.nmap.org
package main

import (
	"fmt"
	"net"
	"sort"
)

const host = "scanme.nmap.org"

func scanner(ports, results chan int) {
	for port := range ports {
		addr := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
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
		go scanner(ports, results)
	}

	// Send ports to scan in a separate goroutine to avoid deadlock. The
	// results gathering loop below has to start. Otherwise the program
	// would get stuck after sending 100 ports to scan. An alternative
	// solution would be to use a buffered channel.
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	var openports []int

	// You can't range over results channel here because the loop would not
	// finish until the channel gets closed.
	for i := 1; i <= 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	sort.Ints(openports)
	fmt.Println(host, openports)
}
