// Package geometry defines simple types for plane geometry. Is shows how to
// declare methods.
//
// Adapted from github.com/adonovan/gopl.io/blob/master/ch6/geometry
package geometry

import "math"

type Point struct{ X, Y float64 }

// Distance is calculated as the length of the hypotenuse [prepona] of a right
// triangle formed by p and q points in a cartesian coordinate system.
func (p Point) Distance(q Point) float64 {
	a := q.X - p.X
	b := q.Y - p.Y
	c := math.Hypot(a, b) // a^2 + b^2 = c^2
	return c
}

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path. (Doesn't conflict the
// Point.Distance method).
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
