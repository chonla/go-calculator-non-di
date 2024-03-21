package calculator

type Calculator struct {
	adder Adder
}

func NewCalculator() Calculator {
	return Calculator{
		adder: NewAdder(),
	}
}

func (c Calculator) Add(x, y int) int {
	return c.adder.Add(x, y)
}
