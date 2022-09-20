// A faster version of lookup.sh that handles errors and can be generalized.
//
// Level: intermediate
// Topics: concurrency, scripting
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

type lookup struct {
	name string

	// Filled in when NS looked up
	err        error
	cloudflare bool
}

func main() {
	var wg sync.WaitGroup
	in := make(chan lookup)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- lookup{name: s.Text()}
		}
		if s.Err() != nil {
			log.Fatalf("Error reading STDIN: %s", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan lookup)

	go func() {
		for l := range out {
			state := "OTHER"
			switch {
			case l.err != nil:
				state = "ERROR"
			case l.cloudflare:
				state = "CLOUDFLARE"
			}

			fmt.Printf("%-10s %s\n", state, l.name)
		}
	}()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for l := range in {
				nss, err := net.LookupNS(l.name)
				if err != nil {
					l.err = err
				} else {
					for _, ns := range nss {
						if strings.HasSuffix(ns.Host, ".ns.cloudflare.com.") {
							l.cloudflare = true
							break
						}
					}
				}
				out <- l
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
