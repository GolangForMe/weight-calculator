package main

import (
	"fmt"

	"github.com/gearmobile/calculator/calc"
	"github.com/gearmobile/calculator/helpers"
)

func main() {
	fmt.Println("___ Калькулятор индекса массы тела ___")

	for {
		weight, height := helpers.GetUserInput()
		imt := calc.Calc(weight, height)
		helpers.PrintImt(imt)
		choise := helpers.CheckUserChoise()

		if choise {
			fmt.Println("Всего хорошего и спасибо за рыбу!")
			break
		}
	}
}
