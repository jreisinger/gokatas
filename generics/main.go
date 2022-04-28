// Show the basics of generics. Based on go.dev/doc/tutorial/generics.
package main

import "fmt"

// Number is called a type constraint.
type Number interface {
	int64 | float64
}

// Sum sums the values of m. It supports both ints and floats. K and V are
// called type parameters. The "comparable" constraint is predeclared in Go.
func Sum[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	ints := map[string]int64{
		"first":  3,
		"second": 14,
	}
	floats := map[string]float64{
		"first":  3.14,
		"second": 31.4,
	}
	fmt.Printf("Sums: %v and %v\n", Sum(ints), Sum(floats))
}
