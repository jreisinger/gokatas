// Decode message encoded with rot13 cipher.
package main

import (
	"io"
	"os"
	"strings"

	"github.com/jreisinger/gokatas/rot13"
)

func main() {
	code := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	d := rot13.Decoder{code}
	io.Copy(os.Stdout, d)
}
