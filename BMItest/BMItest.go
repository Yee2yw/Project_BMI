package BMItest

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
	c := make(chan float64)

	go func() {
		c <- h.calculate()
	}()

	select {
	case num := <-c:
		if num > 24 {
			Println("You should go on a diet!")
			close(c)
		}
		if num < 18.5 {
			Println("You should put on weight!")
			close(c)
		}
	default:
		Println("How fit you are! Try to keep it")
		close(c)
	}
}
