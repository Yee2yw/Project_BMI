package BMI

import (
	. "fmt"
)

type BMI interface {
	calculate() float64
}

type Human struct {
	Height float64
	Weight float64
}

func (h Human) calculate() float64 {
	i := h.Height * h.Height
	return h.Weight / i
}

func Compare(h BMI) {
	switch {
	case h.calculate() > 24:
		Println("You should go on a diet!")
	case h.calculate() < 18.5:
		Println("You should put on weight!")
	default:
		Println("How fit you are! Try to keep it")
	}
}
