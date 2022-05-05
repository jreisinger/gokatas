// Package rot13 implements an io.Reader that reads from an io.Reader, modifying
// the stream by applying the rot13 cipher.
//
// rot13 is a simple sibstitution cipher that rotates all alphabetical
// characters by 13 places. It's decoded by the same algorithm.
//
// Adapted from tour.golang.org/methods/23.
package rot13

import (
	"io"
)

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
	return (b-a+13)%(z-a+1) + a
}

type Rot13 struct {
	Code io.Reader
}

func (r Rot13) Read(p []byte) (int, error) {
	n, err := r.Code.Read(p) // remove Code to get stack overflow error :-)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, err
}
