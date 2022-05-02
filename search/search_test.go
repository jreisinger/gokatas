package search

import (
	"testing"
)

var list = []int{1, 2, 2, 4, 5, 5, 7, 8, 9, 9}

var tests = []struct {
	list []int
	item int
	idx  int
}{
	{list, -1, -1},
	{list, 0, -1},
	{list, 1, 0},
	{list, 2, 1},
	{list, 3, -1},
	{list, 4, 3},
	{list, 5, 4},
	{list, 6, -1},
	{list, 7, 6},
	{list, 8, 7},
	{list, 9, 8},
	{list, 10, -1},
}

func TestLinearsearch(t *testing.T) {
	for _, test := range tests {
		if idx := Linear(test.list, test.item); idx != test.idx {
			t.Errorf("Linear(%v, %d) = %d, want %d", test.list, test.item, idx, test.idx)
		}
	}
}

func TestBinarysearch(t *testing.T) {
	for _, test := range tests {
		if idx := Binary(test.list, test.item); idx != test.idx {
			t.Errorf("Binary(%v, %d) = %d, want %d", test.list, test.item, idx, test.idx)
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

func benchmarkLinear(b *testing.B, size int) {
	list = genList(size)
	for i := 0; i < b.N; i++ {
		Linear(list, size-1)
	}
}
func BenchmarkLinear10(b *testing.B)   { benchmarkLinear(b, 10) }
func BenchmarkLinear100(b *testing.B)  { benchmarkLinear(b, 100) }
func BenchmarkLinear1000(b *testing.B) { benchmarkLinear(b, 1000) }

func benchmarkBinary(b *testing.B, size int) {
	list = genList(size)
	for i := 0; i < b.N; i++ {
		Binary(list, size-1)
	}
}
func BenchmarkBinary10(b *testing.B)   { benchmarkBinary(b, 10) }
func BenchmarkBinary100(b *testing.B)  { benchmarkBinary(b, 100) }
func BenchmarkBinary1000(b *testing.B) { benchmarkBinary(b, 1000) }
