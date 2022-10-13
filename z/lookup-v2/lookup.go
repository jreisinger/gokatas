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
		return
	}
	for _, ns := range nss {
		if strings.HasSuffix(ns.Host, ".cloudflare.com.") {
			l.cloudflare = true
			break
		}
	}
}

func (l *lookupTask) Print() {
	status := "OTHER"
	switch {
	case l.err != nil:
		status = "ERROR"
	case l.cloudflare:
		status = "CLOUDFLARE"
	}
	fmt.Printf("%-10s %s\n", status, l.name)
}

func main() {
	z.Run(lookupFactory{})
}
