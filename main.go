package main

import (
	"fmt"

	"github.com/gearmobile/calculator/calc"
	"github.com/gearmobile/calculator/check"
)

func main() {
	var (
		weight int     = 80
		height float64 = 1.7
	)

	fmt.Println("___ Калькулятор индекса массы тела ___")

	fmt.Print("Введите свой вес в килограммах (80): ")
	fmt.Scan(&weight)

	fmt.Print("Введите свой рост в метрах (1.7): ")
	fmt.Scan(&height)

	imt := calc.Calc(weight, height)

	fmt.Printf("Индекс веса равен %0.2f\n", imt)
	fmt.Println(check.Check(imt))
}
