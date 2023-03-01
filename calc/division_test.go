package calc

import (
	"testing"
)

func TestDivision_Calculate(t *testing.T) {
	type test struct {
		in1  int
		in2  int
		want int
	}

	tests := []test{
		{in1: 0, in2: 1, want: 0},
		{in1: 10, in2: 1, want: 10},
		{in1: 55, in2: 11, want: 5},
		{in1: 10, in2: -2, want: -5},
	}

	for _, i := range tests {
		got := Division{}.Calculate(i.in1, i.in2)
		if got != i.want {
			t.Errorf("got %d, wanted %d", got, i.want)
		}
	}
}
