package geometry

import (
	"fmt"
)

func ExamplePoint_Distance() {
	p := Point{1, 1}
	q := Point{5, 4}
	fmt.Println(p.Distance(q))
	// Output: 5
}

func ExamplePath_Distance() {
	perim := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(perim.Distance())
	// Output: 12
}
