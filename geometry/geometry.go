// Package geometry defines simple types for plane geometry. Adapted from the
// gopl.io ch. 6.1 Method declarations.
package geometry

import "math"

type Point struct{ X, Y float64 }

func (p Point) Distance(q Point) float64 {
	// the length of the hypotenuse [prepona] of a right triangle
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
