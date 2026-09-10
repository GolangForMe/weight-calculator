package main

import (
	"fmt"

	"github.com/gearmobile/calculator/calc"
	"github.com/gearmobile/calculator/check"
)

func getUserInput() (int, float64) {
	var (
		weight int     = 80
		height float64 = 1.7
	)

	fmt.Print("Введите свой вес в килограммах (80): ")
	fmt.Scan(&weight)

	fmt.Print("Введите свой рост в сантиметрах (170): ")
	fmt.Scan(&height)

	return weight, height
}

func printImt(imt float64) {
	fmt.Printf("Индекс веса равен %0.2f\n", imt)
	fmt.Println(check.Check(imt))
}

func main() {
	fmt.Println("___ Калькулятор индекса массы тела ___")
	weight, height := getUserInput()
	imt := calc.Calc(weight, height)
	printImt(imt)
}
