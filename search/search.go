// Package search implements linear and binary search algorithms.
//
//	go test -bench=. # run all tests (if any) and benchmarks
//
// For more see
//   - dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
//   - github.com/jreisinger/homepage/blob/master/notes/cs/big-o-notation.md
//
// Level: intermediate
// Topics: algorithms, testing, benchmarking
package search

// Linear returns the smallest index of item from a list or -1 if not found.
// Linear (or simple) search is an O(n) algorithm.
func Linear(list []int, item int) int {
	for i, n := range list {
		if n == item {
			return i
		}
	}
	return -1
}

// Binary returns the smallest index of item from a *sorted* list or -1 if not
// found. Binary search is an O(log n) algorithm - it has this property that
// tells you how many times you have to divide n with 2 to get 1. For example
// log 16 = 4 => 16 -> 8 -> 4 -> 2 -> 1.
func Binary(list []int, item int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		switch {
		case guess == item:
			return mid
		case guess < item:
			low = mid + 1
		case guess > item:
			high = mid - 1
		}
	}
	return -1
}
