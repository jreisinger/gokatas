// Echosrv is a web server that echoes the path component of the URL. It also
// counts the HTTP requests. NOTE: Handler for each incoming request is run in a
// separate goroutine. To avoid race conditions we need to protect the count
// variable with a mutex.
//
// Level: beginner
// Topics: web server, net/http, locking, sync.Mutex
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

var mu sync.Mutex
var count int

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL path: %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	mu.Unlock()
}
