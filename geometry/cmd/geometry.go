// Geometry calculates distance between two points and the perimeter of a right
// triangle.
package main

import (
	"fmt"

	"github.com/jreisinger/gokatas/geometry"
)

func main() {
	p := geometry.Point{1, 1}
	q := geometry.Point{5, 4}
	fmt.Println(p.Distance(q)) // 5

	perim := geometry.Path{p, {5, 1}, q, {1, 1}}
	fmt.Println(perim.Distance()) // 12
}
