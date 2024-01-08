package search

import (
	"testing"
)

// gentList generates a sorted (necessary for binary search to work) list with n
// elements.
func genList(n int) []int {
	var list []int
	for i := 0; i < n; i++ {
		list = append(list, i)
	}
	return list
}

// benchmarkSearch benchmarks the speed of a search that searches for the
// penultimate item in a list.
func benchmarkSearch(
	search func(list []int, item int) int,
	list []int,
	b *testing.B,
) {
	penultimate := list[len(list)-1]
	for i := 0; i < b.N; i++ {
		search(list, penultimate)
	}
}

var list10 = genList(10)
var list100 = genList(100)
var list1000 = genList(1000)

func BenchmarkLinear10(b *testing.B)   { benchmarkSearch(Linear, list10, b) }
func BenchmarkLinear100(b *testing.B)  { benchmarkSearch(Linear, list100, b) }
func BenchmarkLinear1000(b *testing.B) { benchmarkSearch(Linear, list1000, b) }

func BenchmarkBinary10(b *testing.B)   { benchmarkSearch(Binary, list10, b) }
func BenchmarkBinary100(b *testing.B)  { benchmarkSearch(Binary, list100, b) }
func BenchmarkBinary1000(b *testing.B) { benchmarkSearch(Binary, list1000, b) }
