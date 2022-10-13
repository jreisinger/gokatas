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
	name       string
	err        error
	cloudflare bool
}

func main() {
	in := make(chan lookup)
	var wg sync.WaitGroup

	// Read lines from stdin.
	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- lookup{name: s.Text()}
		}
		if s.Err() != nil {
			log.Fatalf("error reading STDIN: %v", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan lookup)

	// Write status to stdout.
	go func() {
		for l := range out {
			status := "OTHER"
			switch {
			case l.err != nil:
				status = "ERROR"
			case l.cloudflare:
				status = "CLOUDFLARE"
			}

			fmt.Printf("%-10s %s\n", status, l.name)
		}
	}()

	// Do the NS lookups.
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
