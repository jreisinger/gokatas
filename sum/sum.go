// Package sum sums a list of integers using loop and divide-and-conquer
// technique. See reisinge.net/notes/cs/divide-and-conquer for more.
//
//	go test sum/*
//
// Level: intermediate
// Topics: algorithms, testing
package sum

func Loop(s []int) int {
	var sum int
	for _, n := range s {
		sum += n
	}
	return sum
}

func DaC(s []int) int {
	if len(s) == 0 {
		return 0
	}
	return s[0] + DaC(s[1:])
}
