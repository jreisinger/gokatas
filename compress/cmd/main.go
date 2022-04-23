// Compress files concurrently using a wait group.
package main

import (
	"log"
	"os"
	"sync"

	"github.com/jreisinger/gokatas/compress"
)

func main() {
	var wg sync.WaitGroup // a WaitGroup doesn't need to be initialized
	for _, arg := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			if err := compress.Compress(file); err != nil {
				log.Printf("compressing %s: %v", file, err)
			}
		}(arg)
	}
	wg.Wait()
}
