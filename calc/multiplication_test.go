package calc

import (
	"testing"
)

func TestMultiplication_Calculate(t *testing.T) {
	type test struct {
		in1  int
		in2  int
		want int
	}

	tests := []test{
		{in1: 1, in2: 0, want: 0},
		{in1: 0, in2: 1, want: 0},
		{in1: 5, in2: 5, want: 25},
		{in1: 6, in2: -6, want: -36},
	}

	for _, i := range tests {
		got := Multiplication{}.Calculate(i.in1, i.in2)
		if got != i.want {
			t.Errorf("got %d, wanted %d", got, i.want)
		}
	}
}
