// Fetch prints content found at URLs defined as command line arguments. Adapted
// from github.com/adonovan/gopl.io/tree/master/ch1/fetch.
//
// Level: beginner
// Topics: web client, net/http
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		resp, err := http.Get(arg)
		if err != nil {
			log.Print(err)
			continue
		}
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
			continue
		}
		resp.Body.Close()
		fmt.Printf("%s", data)
	}
}
