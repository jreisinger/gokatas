// Calculate square root: given number x, we want to find the number z for
// which z² is most nearly x. Adapted from tour.golang.org/flowcontrol/8.
package main

import (
	"fmt"
	"math"
)

const delta = 1e-15

func sqrt(x float64) float64 {
	z := 1.0 // initial guess
	for math.Abs(z*z-x) > delta {
		z -= (z*z - x) / (2 * z) // Newton's method
	}
	return z
}

func main() {
	fmt.Println(sqrt(100))
}
