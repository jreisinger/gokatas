// Google is a toy search engine using the Go concurrency patterns from the
// boring folder. We start by invoking (fake) Web, Image and Video searches
// serially. We evolve the code until it is concurrent and replicated. See
// https://youtu.be/f6kdp27TYZs?t=1702 for more.
//
// Level: advanced
// Topics: design
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

type Result string

func Google(query string) (results []Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}
