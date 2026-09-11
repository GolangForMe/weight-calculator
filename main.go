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
		imt, err := calc.Calc(weight, height)
		if err != nil {
			fmt.Println(err)
			continue
		}
		helpers.PrintImt(imt)

		if helpers.CheckUserChoise() {
			fmt.Println("Всего хорошего и спасибо за рыбу!")
			break
		}
	}
}
