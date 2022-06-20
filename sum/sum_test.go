package sum

import "testing"

var tests = []struct {
	input []int
	want  int
}{
	{[]int{}, 0},
	{[]int{0}, 0},
	{[]int{1, 3, 2}, 6},
}

func TestLoop(t *testing.T) {
	for _, test := range tests {
		if got := Loop(test.input); got != test.want {
			t.Errorf("Loop(%v) = %d, want %d",
				test.input, got, test.want)
		}
	}
}

func TestDaC(t *testing.T) {
	for _, test := range tests {
		if got := DaC(test.input); got != test.want {
			t.Errorf("DaC(%v) = %d, want %d",
				test.input, got, test.want)
		}
	}
}
