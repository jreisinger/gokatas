// Google is a toy search engine using the Go concurrency patterns from the
// boring package. It's concurrent, replicated and robust. See
// youtube.com/watch?v=f6kdp27TYZs&t=1766s for more.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web1   = fakeSearch("web")
	Web2   = fakeSearch("web")
	Image1 = fakeSearch("image")
	Image2 = fakeSearch("image")
	Video1 = fakeSearch("video")
	Video2 = fakeSearch("video")
)

type Search func(query string) Result

type Result string

func fakeSearch(kind string) Search {
	return func(query string) Result {
		n := rand.Intn(100)
		time.Sleep(time.Duration(n) * time.Millisecond)

		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result) {
	c := make(chan Result)

	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}

	return
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := First("golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"))
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
