package search

import (
	"testing"
)

// Must be sorted for binary search to work.
var list = []int{1, 2, 2, 3, 4, 5, 7, 8, 9}

var tests = []struct {
	list []int
	item int
	idx  int
}{
	{[]int{}, 0, -1},
	{[]int{}, -1, -1},
	{list, -1, -1},
	{list, 0, -1},
	{list, 1, 0},
	{list, 2, 1},
	{list, 6, -1},
	{list, 10, -1},
}

func TestLinear(t *testing.T) {
	for _, test := range tests {
		if idx := Linear(test.list, test.item); idx != test.idx {
			t.Errorf("Linear(%v, %d) = %d, want %d",
				test.list, test.item, idx, test.idx)
		}
	}
}

func TestBinary(t *testing.T) {
	for _, test := range tests {
		if idx := Binary(test.list, test.item); idx != test.idx {
			t.Errorf("Binary(%v, %d) = %d, want %d",
				test.list, test.item, idx, test.idx)
		}
	}
}

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
