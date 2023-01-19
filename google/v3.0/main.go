// V3.0 introduces replication. It means we have multiple search services
// (replicas) for each kind and we take the first result returned by the fastest
// service. This way we dramatically lower the likelihood of discarding results.
// This is a fast and robust program.
//
// Level: advanced
// Topics: concurrency, timeout, replicas
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}

type Result string

func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- First(query, Web1, Web2) }()
	go func() { c <- First(query, Image1, Image2) }()
	go func() { c <- First(query, Video1, Video2) }()

	timeout := time.After(time.Millisecond * 80)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}
	return
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	replica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go replica(i)
	}
	return <-c
}

var (
	Web1   = fakeSearch("web")
	Web2   = fakeSearch("web")
	Image1 = fakeSearch("image")
	Image2 = fakeSearch("image")
	Video1 = fakeSearch("video")
	Video2 = fakeSearch("video")
)

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(time.Millisecond * time.Duration(rand.Intn(100))))
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}
