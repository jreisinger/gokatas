// Fetchall fetches URLs supplied as CLI arguments concurrently and reports
// their times and sizes. Adapted from
// github.com/adonovan/gopl.io/tree/master/ch1/fetchall.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan string)
	for _, arg := range os.Args[1:] {
		go fetch(arg, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.3fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	n, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	// time size url
	ch <- fmt.Sprintf("%.3fs %7d %s", time.Since(start).Seconds(), n, url)
}
