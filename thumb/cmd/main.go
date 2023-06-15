// Thumb generates thumnails for the supplied JPEG pictures.
package main

import (
	"fmt"
	"os"

	"github.com/jreisinger/gokatas/thumb"
)

func main() {
	// thumb.Nail3(os.Args[1:])

	// thumbs, err := thumb.Nail5(os.Args[1:])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(thumbs)

	pics := make(chan string)
	go func() {
		for _, arg := range os.Args[1:] {
			pics <- arg
		}
		close(pics)
	}()
	n := thumb.Nail6(pics)
	fmt.Println(n)
}
