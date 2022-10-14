// Package areader implements io.Reader with a type that emits an infinite
// stream of the ASCII character 'A'. Adapted from tour.golang.org/methods/22.
//
// Level: beginner
// Topics: interfaces, io.Reader
package areader

type Areader struct{}

func (r Areader)Read(slice []byte) error{
	for i:= range slice{
		slice[i] = 'A'
	}
	return nil

}

//definir um tipo Areader
//criar uma funcão Read que recebe um []byte e retorna um erro