package main

import (
	"calculator/calculator"
	"fmt"
)

func main() {
	calc := calculator.NewCalculator()

	fmt.Println(calc.Add(4, 8))
}
