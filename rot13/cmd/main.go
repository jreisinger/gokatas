// Decode message encoded with rot13 cipher.
package main

import (
	"io"
	"os"
	"strings"

	"github.com/jreisinger/gokatas/rot13"
)

func main() {
	ciphertext := "Lbh penpxrq gur pbqr!"
	s := strings.NewReader(ciphertext)
	r := rot13.Reader{R: s}
	io.Copy(os.Stdout, r) // plaintext to STDOUT
}
