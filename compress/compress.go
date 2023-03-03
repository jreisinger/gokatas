// Package compress compresses files. Adapted from the Go in Practice book.
//
// Level: beginner
// Topics: concurrency, sync.WaitGroup
package compress

import (
	"compress/gzip"
	"io"
	"os"
)

func Compress(file string) error {
	in, err := os.Open(file) // dúvida: pq ele abre o file antes de criar o arquivo??
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(file + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gz := gzip.NewWriter(out)
	defer gz.Close()
	_, err = io.Copy(gz, in)
	return err
}
