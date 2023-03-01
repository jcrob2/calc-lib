package calc

type Exponent struct{}

func (Exponent) Calculate(a, b int) int {
	c := 1
	if b == 0 {
		return 1
	}
	for i := 0; i < b; i++ {
		c = c * a
	}
	return c
}
