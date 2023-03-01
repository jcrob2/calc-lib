package calc

import "testing"

func TestExponent_Calculate(t *testing.T) {
	type test struct {
		in1  int
		in2  int
		want int
	}

	tests := []test{
		{in1: 2, in2: 2, want: 4},
		{in1: 1, in2: 6, want: 1},
		{in1: 3, in2: 0, want: 1},
		{in1: 0, in2: 5, want: 0},
	}

	for _, i := range tests {
		got := Exponent{}.Calculate(i.in1, i.in2)
		if got != i.want {
			t.Errorf("got %d, wanted %d", got, i.want)
		}
	}
}
