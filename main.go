package main

import (
	"fmt"

	"github.com/gearmobile/calculator/calc"
	"github.com/gearmobile/calculator/helpers"
)

func main() {
	fmt.Println("___ Калькулятор индекса массы тела ___")
	weight, height := helpers.GetUserInput()
	imt := calc.Calc(weight, height)
	helpers.PrintImt(imt)
}
