// Package areader implements io.Reader with a type that emits an infinite
// stream of the ASCII character 'A'. Adapted from tour.golang.org/methods/22.
package areader

type Areader struct{}

func (r Areader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}
