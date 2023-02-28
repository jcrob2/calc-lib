package calc

type Subtraction struct{}

func (Subtraction) Calculate(a, b int) int {
	return a - b
}
