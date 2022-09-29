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
	return &lookupTask{name: line}
}

type lookupTask struct {
	name       string
	err        error
	cloudflare bool
}

func (l *lookupTask) Process() {
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

func (l *lookupTask) Print() {
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
