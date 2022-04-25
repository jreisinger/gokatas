package search

import (
	"testing"
)

// Tests

type testpair struct {
	list []int
	item int
	idx  int
}

var list = []int{1, 2, 2, 4, 5, 5, 7, 8, 9, 9}

var testpairs = []testpair{
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
	for _, tp := range testpairs {
		idx := Linear(tp.list, tp.item)
		if idx != tp.idx {
			t.Fatalf("Got index %d but wanted %d. Searched %d in %v.", idx, tp.idx, tp.item, tp.list)
		}
	}
}

func TestBinarysearch(t *testing.T) {
	for _, tp := range testpairs {
		idx := Binary(tp.list, tp.item)
		if idx != tp.idx {
			t.Fatalf("Got index %d but wanted %d. Searched %d in %v.", idx, tp.idx, tp.item, tp.list)
		}
	}
}

// Benchmarks (https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)

func genList(n int) []int {
	var list []int
	for i := 0; i < n; i++ {
		list = append(list, i)
	}
	return list
}

var (
	list10   = genList(10)
	list100  = genList(100)
	list1000 = genList(1000)
)

func BenchmarkLinearsearch10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Linear(list10, 9)
	}
}

func BenchmarkLinearsearch100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Linear(list100, 99)
	}
}

func BenchmarkLinearsearch1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Linear(list1000, 999)
	}
}

func BenchmarkBinarysearch10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Binary(list10, 9)
	}
}

func BenchmarkBinarysearch100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Binary(list100, 99)
	}
}

func BenchmarkBinarysearch1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Binary(list1000, 999)
	}
}
