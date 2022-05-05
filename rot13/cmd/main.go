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
	rr := rot13.Rot13{Code: code}
	io.Copy(os.Stdout, rr)
}
