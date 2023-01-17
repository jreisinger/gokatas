// A faster version of lookup-v0.sh that handles errors and can be generalized.
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
	cloudflare bool
	err        error
}

func main() {
	in := make(chan lookup)

	var wg sync.WaitGroup

	// Read lines from stdin and stuff them down the in channel.
	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- lookup{name: s.Text()}
		}
		if s.Err() != nil {
			log.Fatalf("reading STDIN: %v", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan lookup)

	// Read from the in channel, do the NS lookups and stuff the results
	// down the out channel.
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for l := range in {
				nss, err := net.LookupNS(l.name)
				if err != nil {
					l.err = err
				} else {
					for _, ns := range nss {
						if strings.HasSuffix(ns.Host, "cloudflare.com.") {
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

	go func() {
		wg.Wait()
		close(out)
	}()

	// Write the results from the out channel to stdout.
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
}
