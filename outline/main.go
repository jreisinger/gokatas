// Outline prints the structure of the HTML node tree supplied on STDIN.
// Source: https://github.com/adonovan/gopl.io/blob/master/ch5/outline
//
// Usage:
//
//	curl -s https://golang.org | outline
//
// Level: intermediate
// Topics: recursion, golang.org/x/net/html
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push element's tag onto stack
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c) // copy of stack
	}
}
