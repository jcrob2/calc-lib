package calc

import (
	"testing"
)

func TestSubtraction_Calculate(t *testing.T) {
	type test struct {
		in1  int
		in2  int
		want int
	}

	tests := []test{
		{in1: 1, in2: 0, want: 1},
		{in1: 0, in2: 1, want: -1},
		{in1: 555, in2: 111, want: 444},
	}

	for _, i := range tests {
		got := Subtraction{}.Calculate(i.in1, i.in2)
		if got != i.want {
			t.Errorf("got %d, wanted %d", got, i.want)
		}
	}
}
