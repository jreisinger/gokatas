package dac

import "testing"

func TestSum(t *testing.T) {
	testcases := []struct {
		list []int
		sum  int
	}{
		{[]int{}, 0},
		{[]int{0}, 0},
		{[]int{1, 3, 2}, 6},
	}
	for _, tc := range testcases {
		got := Sum(tc.list)
		if got != tc.sum {
			t.Errorf("got %d but wanted %d", got, tc.sum)
		}
	}
}
