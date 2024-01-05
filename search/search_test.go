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

var list10 = genList(10)
var list100 = genList(100)
var list1000 = genList(1000)

func benchmarkLinear(b *testing.B, list []int) {
	for i := 0; i < b.N; i++ {
		Linear(list, list[len(list)-1])
	}
}
func BenchmarkLinear10(b *testing.B)   { benchmarkLinear(b, list10) }
func BenchmarkLinear100(b *testing.B)  { benchmarkLinear(b, list100) }
func BenchmarkLinear1000(b *testing.B) { benchmarkLinear(b, list1000) }

func benchmarkBinary(b *testing.B, list []int) {
	for i := 0; i < b.N; i++ {
		Binary(list, list[len(list)-1])
	}
}
func BenchmarkBinary10(b *testing.B)   { benchmarkBinary(b, list10) }
func BenchmarkBinary100(b *testing.B)  { benchmarkBinary(b, list100) }
func BenchmarkBinary1000(b *testing.B) { benchmarkBinary(b, list1000) }
