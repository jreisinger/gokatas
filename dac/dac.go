// Package dac sums a list of integers using divide-and-conquer technique. See
// reisinge.net/notes/cs/divide-and-conquer for more.
package dac

func Sum(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + Sum(list[1:])
}
