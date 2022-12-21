// Decode message encoded with rot13 cipher.
package main

import (
	"io"
	"os"
	"strings"

	"github.com/jreisinger/gokatas/rot13"
)

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13.Reader{R: s}
	io.Copy(os.Stdout, &r)
}
