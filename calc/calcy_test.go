package calc

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	type test struct {
		in1  int
		in2  int
		want int
	}

	tests := []test{
		{in1: 0, in2: 0, want: 0},
		{in1: 10, in2: 10, want: 20},
		{in1: 1, in2: 2, want: 3},
	}

	for _, i := range tests {
		got := Addition{}.Calculate(i.in1, i.in2)
		if got != i.want {
			t.Errorf("got %d, wanted %d", got, i.want)
		}
	}
}
