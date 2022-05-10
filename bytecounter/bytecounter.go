// Package bytecounter is an implementation of io.Writer that counts bytes.
// Adapted from github.com/adonovan/gopl.io/blob/master/ch7/bytecounter.
package bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to Bytecounter
	return len(p), nil
}
