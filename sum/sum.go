// Package sum sums a list of integers using loop and divide-and-conquer
// technique. See reisinge.net/notes/cs/divide-and-conquer for more.
//	go test sum/*
package sum

func Loop(list []int) int {
	var sum int
	for _, n := range list {
		sum += n
	}
	return sum
}

func DaC(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + DaC(list[1:])
}
