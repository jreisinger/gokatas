// Package rot13 implements an io.Reader that reads from an io.Reader,
// modifying the stream by applying the rot13 cipher. Adapted from
// tour.golang.org/methods/23.
//
// Level: intermediate
// Topics: algorithms, security, io.Reader
package rot13

import (
	"io"
)

// rot13 is a simple sibstitution cipher that rotates all alphabetical
// characters by 13 places. It's decoded by the same algorithm.
func rot13(b byte) byte {
	// Explained here: https://stackoverflow.com/questions/25214008/rot13-and-the-use-of-the-modulo
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

type Decoder struct {
	Code io.Reader
}

func (d Decoder) Read(p []byte) (int, error) {
	n, err := d.Code.Read(p) // remove Code to get stack overflow error :-)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, err
}
