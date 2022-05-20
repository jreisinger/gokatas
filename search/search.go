// Package search implements linear and binary search algorithms.
// 	go test -bench=. search/*	# run tests and all benchmarks
// See reisinge.net/notes/cs/big-o-notation for more.
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
// found. Binary search is an O(log n) algorithm.
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
