// Lookup version that uses z package.
//
// Level: intermediate
// Topics: concurrency, inferfaces, scripting
package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/jreisinger/gokatas/z"
)

type lookupFactory struct{}

func (lookupFactory) Make(line string) z.Task {
	return &lookup{name: line}
}

type lookup struct {
	name       string
	err        error
	cloudflare bool
}

func (l *lookup) Process() {
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
}

func (l *lookup) Print() {
	state := "OTHER"
	switch {
	case l.err != nil:
		state = "ERROR"
	case l.cloudflare:
		state = "CLOUDFLARE"
	}

	fmt.Printf("%-10s %s\n", state, l.name)
}

func main() {
	z.Run(&lookupFactory{})
}
