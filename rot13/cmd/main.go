package main

import (
	"io"
	"os"
	"strings"

	"github.com/jreisinger/gokatas/rot13"
)

func main() {
	r := rot13.Rot13{Code: strings.NewReader("Lbh penpxrq gur pbqr!\n")}
	io.Copy(os.Stdout, &r)
}
