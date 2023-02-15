package main

import "github.com/jcrob2/calc-lib/calc"
import "os"
import "fmt"
import "strconv"

func main() {
	a, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(os.Args[2])

	if err != nil {
		panic(err)
	}

	var calculator calc.Calculator = calc.Addition{}

	result := calculator.Calculate(a, b)
	fmt.Printf("%d + %d = %d", a, b, result)
}
