// Package thumb explores three common concurrency patterns for executing all
// the iterations of a loop in parallel. Adapted from ch. 8.5 of gopl.io.
//
// Topics: concurrency, loop, design
// Level: intermediate
package thumb

import (
	"log"
	"os"
	"sync"

	"github.com/jreisinger/gokatas/thumbnail"
)

// Nail3 makes thumbnails of the specified files in parallel.
func Nail3(filenames []string) {
	ch := make(chan struct{}) // empty struct occupies zero bytes of storage

	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) // NOTE: ignoring potential error
			ch <- struct{}{}
		}(f)
	}

	// Wait for goroutines to complete.
	for range filenames {
		<-ch
	}
}

// Nail5 makes thumbnails of the specified files in parallel. It returns the
// generated file names in an arbitrary order, or an error if any step failed.
func Nail5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	// Buffered channel with sufficient capacity to avoid goroutine leak.
	// This is needed because the Nail5 function might return early when
	// there's an error and the goroutines would block forever.
	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch // no need to drain it because it's buffered
		if it.err != nil {
			return nil, err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

// Nail6 makes thumbnails of each file received from the channel. It returns
// the number of bytes occupied by the files it creates. (Here we can't predict
// the number of iterations because we receive the file names not as a slice but
// over a channel.)
func Nail6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // number of worker goroutines

	for f := range filenames {
		wg.Add(1)

		// worker
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
