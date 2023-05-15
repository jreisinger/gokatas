// Package rot13 implements an io.Reader that reads from an io.Reader, modifying
// the stream by applying the rot13 algorithm to it. Adapted from
// tour.golang.org/methods/23.
//
// Level: intermediate
// Topics: algorithms, security, io.Reader
package rot13

import (
	"io"
)

// rot13 is a simple substitution cipher that rotates all alphabetical
// characters by 13 places. You can use the same algorithm both for encoding and
// decoding (13 is half of 26). Algorithm explanation:
// https://stackoverflow.com/questions/25214008/rot13-and-the-use-of-the-modulo
func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	// return (b-a+13)%26 + a
	return (b-a+13)%(z-a+1) + a
}

type Reader struct {
	R io.Reader
}

func (r Reader) Read(p []byte) (int, error) {
	n, err := r.R.Read(p) // remove R to get stack overflow error :-)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, err
}
