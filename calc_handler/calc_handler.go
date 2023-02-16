package calc_handler

import "C"
import (
	"fmt"
	"github.com/jcrob2/calc-lib/calc"
	"io"
	"strconv"
)

func BridgeConstructor(w io.Writer, c calc.Calculator) Bridge {
	return Bridge{W: w, C: c}
}

type Handler interface {
	Handle(r []string) error
}

type Bridge struct {
	W io.Writer
	C calc.Calculator
}

func (br Bridge) Handle(r []string) error {
	a, err := strconv.Atoi(r[1])

	if err != nil {
		return err
	}

	b, err := strconv.Atoi(r[2])
	if err != nil {
		return err
	}

	result := br.C.Calculate(a, b)
	_, err = fmt.Fprintf(br.W, "%d + %d = %d", a, b, result)
	if err != nil {
		return err
	}

	return nil
}
