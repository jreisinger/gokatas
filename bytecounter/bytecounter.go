// Package bytecounter is an implementation of io.Writer that counts bytes.
// Adapted from github.com/adonovan/gopl.io/blob/master/ch7/bytecounter.
package bytecounter

type Bytecounter int

func (c *Bytecounter) Write(p []byte) (int, error) {
	*c += Bytecounter(len(p)) // convert int to bytecounter
	return len(p), nil
}
