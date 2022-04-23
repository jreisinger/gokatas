// Generics shows the basics of generics. Based on go.dev/doc/tutorial/generics.
package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumIntsOrFloats[K comparable, V Number](m map[K]V) V {
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
	fmt.Printf("Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
	)
}
